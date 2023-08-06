package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/discov"
	"goChat/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	"testing"
)

func TestGetAllByEtcd(t *testing.T) {
	// 获取所有的msg-gateway服务
	msggatewayConfig := discov.EtcdConf{
		Hosts: []string{"0.0.0.0:2379"},
		Key:   "msggateway-rpc",
	}
	services, err := onlinemessagerelayservice.GetAllByEtcd(context.Background(), msggatewayConfig, "msggateway-rpc")

	if err != nil {
		t.Errorf("get all msg gateway service error: %v", err)
		return
	}
	t.Logf("services: %v", services)
}
