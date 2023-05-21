package main

import (
	"flag"
	"fmt"

	"goChat/app/usercenter/cmd/api/internal/config"
	"goChat/app/usercenter/cmd/api/internal/handler"
	"goChat/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the rpcconfig file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting rpcserver at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
