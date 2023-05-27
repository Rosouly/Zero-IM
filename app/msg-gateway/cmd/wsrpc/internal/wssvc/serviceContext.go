package wssvc

import (
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wsconfig"
	"goChat/app/usercenter/cmd/rpc/imuserservice"
)

type ServiceContext struct {
	Config        wsconfig.Config
	ImUserService imuserservice.ImUserService
}

func NewServiceContext(c wsconfig.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
