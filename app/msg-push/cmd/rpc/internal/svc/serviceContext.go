package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/msg-push/cmd/rpc/internal/config"
	"goChat/app/msg-push/cmd/rpc/internal/sdk"
	"goChat/app/usercenter/cmd/rpc/imuserservice"
)

type ServiceContext struct {
	Config        config.Config
	offlinePusher sdk.OfflinePusher
	ImUserService imuserservice.ImUserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ImUserService: imuserservice.NewImUserService(zrpc.MustNewClient(c.ImUserRpc)),
	}
}
func (s *ServiceContext) GetOfflinePusher() sdk.OfflinePusher {
	if s.offlinePusher != nil {
		return s.offlinePusher
	}
	if s.Config.PushType == "jpns" {
		//s.offlinePusher = &push.JPush{s.Config}
	}
	return s.offlinePusher
}
