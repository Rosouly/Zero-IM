package logic

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
	pushpb "goChat/app/msg-push/cmd/rpc/pb"
	"goChat/app/msg-transfer/cmd/persistent/internal/repository"
	"goChat/app/msg-transfer/cmd/persistent/internal/svc"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	"goChat/common/types"
	"goChat/common/utils"
	"goChat/common/utils/statistics"
	"goChat/common/xtrace"
)

var (
	singleMsgSuccessCount uint64
	groupMsgCount         uint64
	superGroupMsgCount    uint64
	singleMsgFailedCount  uint64
)

func init() {
	statistics.NewStatistics(&singleMsgSuccessCount, "msg-transfer-persistent", fmt.Sprintf("%d second singleMsgCount insert to mysql", 300), 300)
	statistics.NewStatistics(&groupMsgCount, "msg-transfer-persistent", fmt.Sprintf("%d second groupMsgCount insert to mysql", 300), 300)
	statistics.NewStatistics(&superGroupMsgCount, "msg-transfer-persistent", fmt.Sprintf("%d second superGroupMsgCount insert to mysql", 300), 300)
}

type MsgTransferPersistentOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewMsgTransferPersistentOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgTransferPersistentOnlineLogic {
	return &MsgTransferPersistentOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *MsgTransferPersistentOnlineLogic) Do(msg []byte, msgKey string) (err error) {
	msgFromMQ := chatpb.MsgDataToMQ{}
	xtrace.StartFuncSpan(l.ctx, "MsgTransferPersistentOnlineLogic.ChatMs2Mongo.UnmarshalMsg", func(ctx context.Context) {
		err = proto.Unmarshal(msg, &msgFromMQ)
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("unmarshal msg failed, err: %v", err)
		return err
	}
	logx.WithContext(l.ctx).Infof("msg: %v", msgFromMQ.String())
	isPersistent := utils.GetSwitchFromOptions(msgFromMQ.MsgData.Options, types.IsPersistent)
	isSenderSync := utils.GetSwitchFromOptions(msgFromMQ.MsgData.Options, types.IsSenderSync)
	switch msgFromMQ.MsgData.SessionType {
	case types.SingleChatType:
		xtrace.StartFuncSpan(l.ctx, "MsgTransferPersistentOnlineLogic.ChatMs2Mongo.SingleChat", func(ctx context.Context) {
			if isPersistent {
				err = l.saveSingleChat(ctx, msgKey, &msgFromMQ)
				if err != nil {
					singleMsgFailedCount++
					l.Logger.Error("single data insert to mysql err ", err.Error(), " ", msgFromMQ.String())
					return
				}
				singleMsgSuccessCount++
			}
			if !isSenderSync && msgKey == msgFromMQ.MsgData.SendID {
			} else {
				go l.sendMessageToPush(ctx, &msgFromMQ, msgKey)
			}
		})
	case types.GroupChatType:
		xtrace.StartFuncSpan(l.ctx, "MsgTransferPersistentOnlineLogic.ChatMs2Mongo.SuperGroupChat", func(ctx context.Context) {
			if isPersistent {
				err = l.saveSuperGroupChat(ctx, msgFromMQ.MsgData.GroupID, &msgFromMQ)
				if err != nil {
					l.Logger.Error("super group data insert to mysql err ", msgFromMQ.String(), " GroupID ", msgFromMQ.MsgData.GroupID, " ", err.Error())
					return
				}
				superGroupMsgCount++
			}
		})
	case types.NotificationChatType:
		xtrace.StartFuncSpan(l.ctx, "MsgTransferPersistentOnlineLogic.ChatMs2Mongo.NotificationChat", func(ctx context.Context) {
			if isPersistent {
				err = l.saveNotificationChat(ctx, msgKey, &msgFromMQ)
				if err != nil {
					l.Logger.Error("single data insert to mysql err ", err.Error(), " ", msgFromMQ.String())
					return
				}
			}
		})
	default:
		l.Logger.Error("SessionType error ", msgFromMQ.String())
		return
	}
	return
}

func (l *MsgTransferPersistentOnlineLogic) sendMessageToPush(ctx context.Context, message *chatpb.MsgDataToMQ, pushToUserID string) {
	logx.WithContext(ctx).Info("msg_transfer send message to push", "message", message.String())
	rpcPushMsg := pushpb.PushMsgReq{MsgData: message.MsgData, PushToUserID: pushToUserID}
	_, err := l.svcCtx.MsgPush.PushMsg(ctx, &rpcPushMsg)
	if err != nil {
		logx.WithContext(ctx).Error("rpc send failed", "push data", rpcPushMsg.String(), "err", err.Error())
		mqPushMsg := chatpb.PushMsgDataToMQ{MsgData: message.MsgData, PushToUserID: pushToUserID, OperationID: xtrace.TraceIdFromContext(l.ctx)}
		pid, offset, err := l.svcCtx.SinglePushProducer.SendMessage(ctx, &mqPushMsg)
		if err != nil {
			logx.WithContext(ctx).Error("kafka send failed", mqPushMsg.OperationID, "send data", mqPushMsg.String(), "pid", pid, "offset", offset, "err", err.Error())
		}
	} else {
		logx.WithContext(ctx).Info("rpc send success", "push data", rpcPushMsg.String())
	}
}
