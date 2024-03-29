package xkafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type MConsumerGroup struct {
	sarama.ConsumerGroup
	groupID string
	topics  []string
}

type MConsumerGroupConfig struct {
	KafkaVersion   sarama.KafkaVersion
	OffsetsInitial int64
	IsReturnErr    bool
}

func NewMConsumerGroup(consumerConfig *MConsumerGroupConfig, topics, addr []string, groupID string) *MConsumerGroup {
	config := sarama.NewConfig()
	config.Version = consumerConfig.KafkaVersion
	config.Consumer.Offsets.Initial = consumerConfig.OffsetsInitial
	config.Consumer.Return.Errors = consumerConfig.IsReturnErr
	client, err := sarama.NewClient(addr, config)
	if err != nil {
		panic(err.Error())
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		panic(err.Error())
	}
	logx.Infof("consumer group:%s , topics:%v , addr:%v", groupID, topics, addr)
	return &MConsumerGroup{
		consumerGroup,
		groupID,
		topics,
	}
}
func (mc *MConsumerGroup) RegisterHandleAndConsumer(handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		err := mc.ConsumerGroup.Consume(ctx, mc.topics, handler)
		logx.Info("consumer err: ", err)
		if err != nil {
			panic(err.Error())
		}

	}
}
