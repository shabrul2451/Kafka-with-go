package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// broker address
	brokers := []string{"localhost:9092"}

	// producer configuration
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// new producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()
	headers := []sarama.RecordHeader{
		{Key: []byte("header_key"), Value: []byte("header_value")},
	}
	message := &sarama.ProducerMessage{
		Topic:     "test_topic", // kafka topic
		Key:       sarama.StringEncoder("demo"),
		Value:     sarama.StringEncoder("Hello from Go Kafka Producer with custom header!"),
		Headers:   headers,
		Timestamp: time.Now().UTC(),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d with offset %d\n", partition, offset)
}
