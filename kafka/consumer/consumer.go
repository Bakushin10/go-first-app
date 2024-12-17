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

	subscribedTopics := []string{"topic-subscribed-1", "topic-subscribed-2", "topic-subscribed-3"}

	partitionConsumers := make(map[string]sarama.PartitionConsumer)

	// Subscribe to all topics and create partition consumers
	for _, topic := range subscribedTopics {
		// Start consuming partition 0 for each topic
		partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Error starting partition consumer for topic %s: %v\n", topic, err)
			continue
		}
		partitionConsumers[topic] = partitionConsumer
		fmt.Printf("Started consuming from topic: %s\n", topic)
		defer partitionConsumer.Close()
	}

	fmt.Println("Consumer started. Waiting for messages...")

	messageChannel := make(chan *sarama.ConsumerMessage)

	// Start a goroutine for each partition consumer to read messages concurrently
	for topic, partitionConsumer := range partitionConsumers {
		go func(topic string, partitionConsumer sarama.PartitionConsumer) {
			for msg := range partitionConsumer.Messages() {
				messageChannel <- msg
			}
		}(topic, partitionConsumer)
	}

	// Consume messages from the channel
	for msg := range messageChannel {
		fmt.Printf("Received message from topic %s: %s\n", msg.Topic, string(msg.Value))
	}
}
