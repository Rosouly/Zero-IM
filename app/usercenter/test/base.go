package rpc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/usercenter/cmd/rpc/imuserservice"
	"goChat/app/usercenter/cmd/rpc/usercenterservice"
)

var (
	imuserConf = zrpc.RpcClientConf{
		Endpoints: []string{
			"0.0.0.0:8082",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	imUserService  = imuserservice.NewImUserService(zrpc.MustNewClient(imuserConf))
	usercenterConf = zrpc.RpcClientConf{
		Endpoints: []string{
			"0.0.0.0:8888",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  5,
	}
	usercenterService = usercenterservice.NewUsercenterService(zrpc.MustNewClient(usercenterConf))
)
