package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:19091",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	}

	if err = c.SubscribeTopics([]string{"myTopic"}, nil); err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
