package main

import (
	"flag"
	"fmt"

	"goChat/app/usercenter/cmd/rpc/internal/config"
	"goChat/app/usercenter/cmd/rpc/internal/server"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var usercenterConfigFile = flag.String("f", "etc/usercenter.yaml", "the rpcconfig file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*usercenterConfigFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServiceServer(grpcServer, server.NewUsercenterServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc rpcserver at %s...\n", c.ListenOn)
	s.Start()
}
