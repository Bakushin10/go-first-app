package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// Create a new Sarama consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Error creating consumer: ", err)
	}
	defer consumer.Close()

	// Subscribe to a topic
	topic := "test-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error starting partition consumer: ", err)
	}
	defer partitionConsumer.Close()

	// Consume messages
	fmt.Println("Consumer started. Waiting for messages...")
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
