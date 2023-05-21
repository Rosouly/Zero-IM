package rpcsvc

import "goChat/app/msg-gateway/cmd/wsrpc/internal/rpcconfig"

type ServiceContext struct {
	Config rpcconfig.Config
}

func NewServiceContext(c rpcconfig.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
