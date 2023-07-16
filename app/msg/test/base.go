package rpc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/msg/cmd/rpc/chat"
)

var (
	msgConf = zrpc.RpcClientConf{
		Endpoints: []string{
			"0.0.0.0:8088",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	msgService = chat.NewChat(zrpc.MustNewClient(msgConf))
)
