package logic

import (
	"github.com/zeromicro/go-zero/core/logx"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	imuserpb "goChat/app/usercenter/cmd/rpc/pb"
	"goChat/common/types"
	"goChat/common/utils"
	timeUtils "goChat/common/utils/time"
	"goChat/common/xorm/global"
)

func (l *SendMsgLogic) encapsulateMsgData(msg *chatpb.MsgData) {
	// 生成消息id
	msg.ServerMsgID = global.GetID()
	msg.ServerTime = timeUtils.GetCurrentTimestampByMill()
	switch msg.ContentType {
	case types.Text:
		fallthrough
	case types.Picture:
		fallthrough
	case types.Voice:
		fallthrough
	case types.Video:
		fallthrough
	case types.File:
		fallthrough
	case types.Merger:
		fallthrough
	case types.Card:
		fallthrough
	case types.Location:
		fallthrough
	case types.Custom:
		fallthrough
	case types.Quote:
		utils.SetSwitchFromOptions(msg.Options, types.IsConversationUpdate, true)
		utils.SetSwitchFromOptions(msg.Options, types.IsUnreadCount, true)
		utils.SetSwitchFromOptions(msg.Options, types.IsSenderSync, true)
	case types.Revoke:
		utils.SetSwitchFromOptions(msg.Options, types.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsOfflinePush, false)
	case types.HasReadReceipt:
		l.Logger.Info("", "this is a test start ", msg, msg.Options)
		utils.SetSwitchFromOptions(msg.Options, types.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsSenderConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsOfflinePush, false)
		l.Logger.Info("", "this is a test end ", msg, msg.Options)
	case types.Typing:
		utils.SetSwitchFromOptions(msg.Options, types.IsHistory, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsPersistent, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsSenderSync, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsSenderConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, types.IsOfflinePush, false)

	}
}

func (l *SendMsgLogic) modifyMessageByUserMessageReceiveOpt(userID, sourceID string, sessionType int, pb *chatpb.SendMsgReq) bool {
	logx.WithContext(l.ctx).Info("modifyMessageByUserMessageReceiveOpt")
	conversationID := types.GetConversationIDBySessionType(sourceID, sessionType)
	// 用户设置了消息接收选项
	req := &imuserpb.GetSingleConversationRecvMsgOptsReq{
		UserID:         userID,
		ConversationID: conversationID,
	}
	resp, err := l.svcCtx.ImUser.GetSingleConversationRecvMsgOpts(l.ctx, req)
	if err != nil {
		logx.WithContext(l.ctx).Error("GetSingleConversationMsgOpt from redis err ", conversationID, " ", pb.String(), " ", err.Error())
		return true
	} else if resp.CommonResp.ErrCode != 0 {
		return true
	} else {
		switch resp.Opts {
		case imuserpb.RecvMsgOpt_NotReceiveMessage:
			return false
		case imuserpb.RecvMsgOpt_ReceiveNotNotifyMessage:
			if pb.MsgData.Options == nil {
				pb.MsgData.Options = make(map[string]bool, 10)
			}
			utils.SetSwitchFromOptions(pb.MsgData.Options, types.IsOfflinePush, false)
		}
		return true
	}
}
