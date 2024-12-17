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
	topic := []string{"topic-subscribed-1", "topic-subscribed-2", "topic-subscribed-3", "topic-not-subscribed"}
	messages := make(map[string]string)
	for _, topic := range topic {
		messages[topic] = fmt.Sprintf("message from %s", topic)
	}

	i := 0
	for {
		selected_topic := topic[i%len(topic)]
		message := messages[selected_topic]

		msg := &sarama.ProducerMessage{
			Topic: selected_topic,
			Value: sarama.StringEncoder(message),
		}

		// Send message to Kafka
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("Error sending message: ", err)
		} else {
			fmt.Println("Message sent: topic:", msg.Topic)
		}
		i++
		// Sleep for a while before sending next message
		time.Sleep(2 * time.Second)
	}
}
