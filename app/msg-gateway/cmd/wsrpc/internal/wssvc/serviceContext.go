package wssvc

import (
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wsconfig"
)

type ServiceContext struct {
	Config wsconfig.Config
}

func NewServiceContext(c wsconfig.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
