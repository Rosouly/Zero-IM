package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/usercenter/cmd/api/internal/config"
	"goChat/app/usercenter/cmd/api/internal/middleware"
	"goChat/app/usercenter/cmd/rpc/usercenterservice"
)

type ServiceContext struct {
	Config        config.Config
	JwqAuth       rest.Middleware
	UserCenterRpc usercenterservice.UsercenterService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		JwqAuth:       middleware.NewJwqAuthMiddleware().Handle,
		UserCenterRpc: usercenterservice.NewUsercenterService(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
