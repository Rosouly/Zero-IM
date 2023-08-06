package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Shopify/sarama"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func TestComsumer(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // Use the correct Kafka version here
	config.Consumer.Return.Errors = true
	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "test-group", config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}

	ctx := context.Background()
	for {
		topics := []string{"im_msg_online"}
		handler := exampleConsumerGroupHandler{}

		err := group.Consume(ctx, topics, handler)
		if err != nil {
			log.Fatalf("Error from consumer: %v", err)
		}
	}
}
