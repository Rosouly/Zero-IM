package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/msg/cmd/rpc/internal/svc"
	"goChat/app/msg/cmd/rpc/pb"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	imuserpb "goChat/app/usercenter/cmd/rpc/pb"
	"goChat/common/types"
	"goChat/common/utils"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func returnMsg(replay *chatpb.SendMsgResp, pb *chatpb.SendMsgReq, errCode int32, errMsg, serverMsgID string, sendTime int64) (*chatpb.SendMsgResp, error) {
	replay.ErrCode = errCode
	replay.ErrMsg = errMsg
	replay.ServerMsgID = serverMsgID
	replay.ClientMsgID = pb.MsgData.ClientMsgID
	replay.ServerTime = sendTime
	return replay, nil
}

func (l *SendMsgLogic) SendMsg(pb *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	logx.WithContext(l.ctx).Info("this is a test SendMsgReq ", pb)
	replay := chatpb.SendMsgResp{}
	flag, errCode, errMsg := l.userRelationshipVerification(pb)
	fmt.Println("userRelationshipVerification end")
	if !flag {
		return returnMsg(&replay, pb, errCode, errMsg, "", 0)
	}
	//if !utils.VerifyToken(pb.Token, pb.SendID) {
	//	return returnMsg(&replay, pb, http.StatusUnauthorized, "token validate err,not authorized", "", 0)
	l.encapsulateMsgData(pb.MsgData)
	fmt.Println("encapsulateMsgData end")
	logx.WithContext(l.ctx).Info("this is a test MsgData ", pb.MsgData)
	msgToMQSingle := chatpb.MsgDataToMQ{Token: pb.Token, MsgData: pb.MsgData}
	fmt.Println("MQSingle end")
	//options := utils.JsonStringToMap(pbData.Options)
	isHistory := utils.GetSwitchFromOptions(pb.MsgData.Options, types.IsHistory)
	mReq := MsgCallBackReq{
		SendID:      pb.MsgData.SendID,
		RecvID:      pb.MsgData.RecvID,
		Content:     string(pb.MsgData.Content),
		SendTime:    pb.MsgData.ServerTime,
		MsgFrom:     pb.MsgData.MsgFrom,
		ContentType: pb.MsgData.ContentType,
		SessionType: pb.MsgData.SessionType,
		PlatformID:  pb.MsgData.SenderPlatformID,
		MsgID:       pb.MsgData.ClientMsgID,
	}
	if !isHistory {
		mReq.IsOnlineOnly = true
	}

	switch pb.MsgData.SessionType {
	case types.SingleChatType:
		fmt.Println("SingleChatType start")
		//isSend := l.modifyMessageByUserMessageReceiveOpt(pb.MsgData.RecvID, pb.MsgData.SendID, types.SingleChatType, pb)
		//if isSend {
		msgToMQSingle.MsgData = pb.MsgData
		logx.WithContext(l.ctx).Info(msgToMQSingle.String())
		// 发送消息，token，接收方id，以及在线状态给kafka
		err1 := l.sendMsgToKafka(&msgToMQSingle, msgToMQSingle.MsgData.RecvID, types.OnlineStatus)
		if err1 != nil {
			logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:RecvID ", msgToMQSingle.MsgData.RecvID, msgToMQSingle.String())
			return returnMsg(&replay, pb, 201, "kafka send msg err ", "", 0)
		}
		//}
		if msgToMQSingle.MsgData.SendID != msgToMQSingle.MsgData.RecvID { //Filter messages sent to yourself
			err2 := l.sendMsgToKafka(&msgToMQSingle, msgToMQSingle.MsgData.SendID, types.OnlineStatus)
			if err2 != nil {
				logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:SendID ", msgToMQSingle.MsgData.SendID, msgToMQSingle.String())
				return returnMsg(&replay, pb, 201, "kafka send msg err ", "", 0)
			}
		}
		return returnMsg(&replay, pb, 0, "", msgToMQSingle.MsgData.ServerMsgID, msgToMQSingle.MsgData.ServerTime)

	case types.GroupChatType:
		// 读扩散
		return returnMsg(&replay, pb, 0, "", "", 0)
	case types.NotificationChatType:
		return returnMsg(&replay, pb, 0, "", "", 0)
	default:
		return returnMsg(&replay, pb, 203, "unkonwn sessionType", "", 0)
	}
}

func (l *SendMsgLogic) userRelationshipVerification(data *chatpb.SendMsgReq) (bool, int32, string) {
	logx.WithContext(l.ctx).Info("userRelationshipVerification start")
	if data.MsgData.SessionType == types.GroupChatType {
		return true, 0, ""
	}
	// 是不是拉黑了
	ifInBlackResp, err := l.svcCtx.ImUser.IfAInBBlacklist(l.ctx, &imuserpb.IfAInBBlacklistReq{
		AUserID: data.MsgData.SendID,
		BUserID: data.MsgData.RecvID,
	})
	logx.WithContext(l.ctx).Info("ifInBlackResp", ifInBlackResp)
	if err != nil {
		logx.WithContext(l.ctx).Error("GetBlackIDListFromCache rpc call failed ", err.Error())
	} else {
		if ifInBlackResp.CommonResp.ErrCode != 0 {
			logx.WithContext(l.ctx).Error("GetBlackIDListFromCache rpc logic call failed ", ifInBlackResp.String())
		} else {
			if ifInBlackResp.IsInBlacklist {
				return false, 600, "in black list"
			}
		}
	}
	if l.svcCtx.Config.MessageVerify.FriendVerify {
		needFriend := true
		switch data.MsgData.ContentType {
		case types.NotificationUser2User:
			needFriend = false
		}
		if !needFriend {
			return true, 0, ""
		}
		// 是不是好友
		ifInFriendResp, err := l.svcCtx.ImUser.IfAInBFriendList(l.ctx, &imuserpb.IfAInBFriendListReq{
			AUserID: data.MsgData.SendID,
			BUserID: data.MsgData.RecvID,
		})
		if err != nil {
			logx.WithContext(l.ctx).Error("GetFriendIDListFromCache rpc call failed ", err.Error())
		} else {
			if ifInFriendResp.CommonResp.ErrCode != 0 {
				logx.WithContext(l.ctx).Error("GetFriendIDListFromCache rpc logic call failed ", ifInFriendResp.String())
			} else {
				if !ifInFriendResp.IsInFriendList {
					return false, 601, "not friend"
				}
			}
		}
		return true, 0, ""
	} else {
		return true, 0, ""
	}
}
