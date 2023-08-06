package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/msg-push/cmd/rpc/msgpushservice"
	"goChat/app/msg-transfer/cmd/persistent/internal/config"
	"goChat/common/xkafka"
)

type ServiceContext struct {
	Config             config.Config
	SinglePushProducer *xkafka.Producer // 单聊的topic
	MsgPush            msgpushservice.MsgPushService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		SinglePushProducer: xkafka.MustNewProducer(c.Kafka.SinglePush),
		MsgPush:            msgpushservice.NewMsgPushService(zrpc.MustNewClient(c.MsgPushRpc)),
	}
}
