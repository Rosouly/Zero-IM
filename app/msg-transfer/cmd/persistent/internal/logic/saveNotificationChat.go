package logic

import (
	"context"
	chatpb "goChat/app/msg/cmd/rpc/pb"
)

func (l *MsgTransferPersistentOnlineLogic) saveNotificationChat(ctx context.Context, key string, c *chatpb.MsgDataToMQ) error {
	// 不存储通知消息
	return nil
}
