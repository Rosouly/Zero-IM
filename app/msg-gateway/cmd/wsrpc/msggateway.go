package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/handler"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcconfig"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcserver"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wsconfig"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wssvc"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
)

var wsConfigFile = flag.String("w", "etc/msggateway-ws.yaml", "ws rpcconfig file")
var rpcConfigFile = flag.String("r", "etc/msggateway-rpc.yaml", "rpc rpcconfig file")

func ws() {
	flag.Parse()

	var c wsconfig.Config
	conf.MustLoad(*wsConfigFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := wssvc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting rpcserver at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func rpc() {
	flag.Parse()

	var c rpcconfig.Config
	conf.MustLoad(*rpcConfigFile, &c)
	ctx := rpcsvc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOnlineMessageRelayServiceServer(grpcServer, rpcserver.NewOnlineMessageRelayServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc rpcserver at %s...\n", c.ListenOn)
	s.Start()
}

func main() {
	go ws()
	logx.Info("ws 启动成功 等待1秒启动 rpc")
	time.Sleep(time.Second)
	rpc()
}
