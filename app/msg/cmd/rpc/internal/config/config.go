package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/common/xkafka"
)

type Config struct {
	zrpc.RpcServerConf
	MessageVerify MessageVerifyConfig
	ImUserRpc     zrpc.RpcClientConf
	Kafka         KafkaConfig
}

type MessageVerifyConfig struct {
	FriendVerify bool // 只有好友才能发送消息
}

type KafkaConfig struct {
	Online  xkafka.ProducerConfig
	Offline xkafka.ProducerConfig
}
