package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	// Create a new Sarama producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Error creating producer: ", err)
	}
	defer producer.Close()

	// Produce messages
	topic := "test-topic"
	message := "Hello from Producer!"

	for {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		// Send message to Kafka
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("Error sending message: ", err)
		} else {
			fmt.Println("Message sent:", message)
		}

		// Sleep for a while before sending next message
		time.Sleep(2 * time.Second)
	}
}
