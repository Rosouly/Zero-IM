package logic

import (
	"context"
	"goChat/app/msg-transfer/model"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	xormerr "goChat/common/xorm/err"
	"goChat/common/xtrace"
)

func (l *MsgTransferPersistentOnlineLogic) saveSingleChat(ctx context.Context, key string, c *chatpb.MsgDataToMQ) error {
	if c == nil || c.MsgData == nil || c.MsgData.ServerMsgID == "" {
		l.Error("saveSingleChat error, c is nil or c.MsgData is nil")
		return nil
	}
	if c.MsgData.SendID == c.MsgData.RecvID {
		// 自己给自己发消息 不存储
		return nil
	}
	chat := &model.MysqlSingleChat{
		ServerMsgID:      c.MsgData.ServerMsgID,
		SendID:           c.MsgData.SendID,
		RecvID:           c.MsgData.RecvID,
		GroupID:          c.MsgData.GroupID,
		ClientMsgID:      c.MsgData.ClientMsgID,
		SenderPlatformID: c.MsgData.SenderPlatformID,
		SenderNickname:   c.MsgData.SenderNickname,
		SenderFaceURL:    c.MsgData.SenderFaceURL,
		SessionType:      c.MsgData.SessionType,
		MsgFrom:          c.MsgData.MsgFrom,
		ContentType:      c.MsgData.ContentType,
		Content:          string(c.MsgData.Content),
		Seq:              c.MsgData.Seq,
		SendTime:         c.MsgData.ServerTime,
		CreateTime:       c.MsgData.ClientTime,
		OfflinePushInfo:  model.NewOfflinePushInfo(c.MsgData.OfflinePushInfo),
		AtUserIDList:     c.MsgData.AtUserIDList,
		Options:          c.MsgData.Options,
	}
	var err error
	xtrace.StartFuncSpan(ctx, "saveSingleChat", func(ctx context.Context) {
		err = chat.Insert(l.rep.Mysql)
		if err != nil {
			if xormerr.DuplicateError(err) {
				err = nil
			}
		}
	})
	if err != nil {
		l.Error("saveSingleChat error", err)
		return err
	}
	return nil
}
