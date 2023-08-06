package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/common/xkafka"
	"goChat/common/xorm/global"
)

type Config struct {
	service.ServiceConf
	Kafka       KafkaConfig
	MysqlConfig global.MysqlConfig
	MsgPushRpc  zrpc.RpcClientConf
}
type KafkaConfigOnline struct {
	xkafka.ProducerConfig
	MsgToMysqlGroupID string
}
type KafkaConfig struct {
	Online KafkaConfigOnline
	//Offline        xkafka.ProducerConfig
	SinglePush xkafka.ProducerConfig
	//SuperGroupPush xkafka.ProducerConfig
}
