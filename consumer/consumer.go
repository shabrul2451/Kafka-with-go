package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker addresses
	brokers := []string{"localhost:9092"}

	// Create a new consumer
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	// Define the topic and partition to consume from
	topic := "test_topic"
	partition := 0

	// Create a partition consumer
	partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for partition: %v", err)
	}
	defer partitionConsumer.Close()

	// Consume messages
	fmt.Println("Consumer is listening for messages...")
	for message := range partitionConsumer.Messages() {
		// checking header
		fmt.Println(string(message.Headers[0].Key), "header")

		// checking partition
		fmt.Println(message.Partition, "partition")

		// checking message
		fmt.Printf("Consumed message from topic: %s, partition: %d, offset: %d, key: %s, value: %s\n",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}
}
