package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/msg/cmd/rpc/internal/config"
	"goChat/app/usercenter/cmd/rpc/imuserservice"
	"goChat/common/xkafka"
)

type ServiceContext struct {
	Config          config.Config
	ImUser          imuserservice.ImUserService
	OnlineProducer  *xkafka.Producer
	OfflineProducer *xkafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ImUser:          imuserservice.NewImUserService(zrpc.MustNewClient(c.ImUserRpc)),
		OnlineProducer:  xkafka.MustNewProducer(c.Kafka.Online),
		OfflineProducer: xkafka.MustNewProducer(c.Kafka.Offline),
	}
}
