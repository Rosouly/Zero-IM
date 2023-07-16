package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"goChat/common/xkafka"
	"goChat/common/xorm/global"
)

type Config struct {
	service.ServiceConf
	Kafka       KafkaConfig
	MysqlConfig global.MysqlConfig
}
type KafkaConfigOnline struct {
	xkafka.ProducerConfig
	MsgToMysqlGroupID string
}
type KafkaConfig struct {
	Online KafkaConfigOnline
}
