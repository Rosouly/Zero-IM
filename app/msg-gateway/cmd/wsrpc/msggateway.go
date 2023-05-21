package main

import (
	"flag"
	"fmt"

	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcconfig"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcserver"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/msggateway.yaml", "the rpcconfig file")

func main() {
	flag.Parse()

	var c rpcconfig.Config
	conf.MustLoad(*configFile, &c)
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
