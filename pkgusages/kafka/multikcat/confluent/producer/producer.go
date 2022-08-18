package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaBrokers = []string{"0.0.0.0:19091"}
	KafkaTopic   = "myTopic"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBrokers[0]})
	if err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v, msg is %v", ev.TopicPartition, string(ev.Value))
				}
			}
		}
	}()

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	produceMsg(producer, signals)
}

func produceMsg(producer *kafka.Producer, signals chan os.Signal) {
	for {
		time.Sleep(time.Second)
		word := "Welcome to the Confluent Kafka Golang client"
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &KafkaTopic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}

		select {
		case producer.ProduceChannel() <- msg:
			log.Println("New Message produced")
		case <-signals:
			producer.Flush(15 * 1000) // Wait for message deliveries, shutting down gracefully
			return
		}
	}

	// The way after sending all messages shut down by itself.
	// for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	// 	_ = producer.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &KafkaTopic, Partition: kafka.PartitionAny},
	// 		Value:          []byte(word),
	// 	}, nil)
	// }

	// Wait for message deliveries before shutting down
	// producer.Flush(15 * 1000)
}
