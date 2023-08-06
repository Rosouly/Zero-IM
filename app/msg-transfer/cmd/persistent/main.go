package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/msg-transfer/cmd/persistent/internal/config"
	"goChat/app/msg-transfer/cmd/persistent/internal/server"
	"goChat/app/msg-transfer/cmd/persistent/internal/svc"
	"goChat/common/xconf"
)

var configFile = flag.String("f", "etc/persistent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*configFile, &c)
	logx.Info("config: %v", c)
	err := c.ServiceConf.SetUp()
	if err != nil {
		panic(err)
	}
	s := server.NewMsgTransferPersistentServer(svc.NewServiceContext(c))
	s.Start()
}
