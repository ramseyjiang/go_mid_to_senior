package consumer

import (
	"context"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers    = []string{"localhost:9093"}
	kafkaTopics     = []string{"sarama_topic"}
	consumerGroupID = "sarama_consumer"
)

func StartConsumer() {
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	config.Version = sarama.MaxVersion

	// Start with a client
	client, err := sarama.NewClient(kafkaBrokers, config)
	if err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(consumerGroupID, client)
	if err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	}
	defer func() { _ = group.Close() }()
	log.Println("Consumer up and running")

	// Track errors
	go func() {
		for err := range group.Errors() {
			log.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := GroupHandler{}

		err := group.Consume(ctx, kafkaTopics, handler)
		if err != nil {
			log.Fatalf("fatal error is:%v\n", err)
		}
	}
}
