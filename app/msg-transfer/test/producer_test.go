package test

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	brokers := []string{"localhost:9092"}
	topic := "im_msg_online"
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
		os.Exit(1)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello, World!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
		os.Exit(1)
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	time.Sleep(5 * time.Second)
}
