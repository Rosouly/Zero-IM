package rpc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/usercenter/cmd/rpc/imuserservice"
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
	imUserService = imuserservice.NewImUserService(zrpc.MustNewClient(imuserConf))
)
