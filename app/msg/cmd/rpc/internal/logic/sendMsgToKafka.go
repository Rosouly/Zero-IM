package logic

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	"goChat/common/types"
	"goChat/common/xtrace"
)

func (l *SendMsgLogic) sendMsgToKafka(m *chatpb.MsgDataToMQ, key string, status string) error {
	logx.WithContext(l.ctx).Info("sendMsgToKafka start")
	m.OperationID = xtrace.TraceIdFromContext(l.ctx)
	switch status {
	case types.OnlineStatus:
		logx.WithContext(l.ctx).Info("types.OnlineStatus")
		pid, offset, err := l.svcCtx.OnlineProducer.SendMessage(l.ctx, m, key)
		if err != nil {
			l.Logger.Error(m.OperationID, "kafka send failed ", "send data ", m.String(), "pid ", pid, "offset ", offset, "err ", err.Error(), "key ", key, status)
		}
		return err
	case types.OfflineStatus:
		pid, offset, err := l.svcCtx.OfflineProducer.SendMessage(l.ctx, m, key)
		if err != nil {
			l.Logger.Error(m.OperationID, "kafka send failed ", "send data ", m.String(), "pid ", pid, "offset ", offset, "err ", err.Error(), "key ", key, status)
		}
		return err
	}
	return errors.New("status error")
}
