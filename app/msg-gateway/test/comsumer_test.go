package test

import (
	"github.com/Shopify/sarama"
	"testing"
	"time"
)

func TestComsumer(t *testing.T) {
	brokers := []string{"localhost:9092"}
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, consumerConfig)
	if err != nil {
		t.Errorf("Error creating consumer: %v",
			err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Errorf("Error closing consumer: %v", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("im_msg_online", 0, sarama.OffsetNewest)
	if err != nil {
		t.Errorf("Error consuming partition: %v", err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			t.Errorf("Error closing partition consumer: %v", err)
		}
	}()

	select {
	case message := <-partitionConsumer.Messages():
		t.Log("message: ", message)
	case <-time.After(time.Second * 10):
		t.Error("time out")
	}

}
