package main

import (
	"context"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers    = []string{"0.0.0.0:19091"}
	kafkaTopics     = []string{"myTopic"}
	consumerGroupID = "myGroup"
)

func main() {
	// setup sarama log
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)

	// Init config, specify appropriate version
	client, err := setupClient()
	defer func() { _ = client.Close() }()

	consume(err, client)
}

func setupClient() (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion

	// Start with a client
	client, err := sarama.NewClient(kafkaBrokers, config)
	if err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	}
	return client, err
}

func consume(err error, client sarama.Client) {
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

// GroupHandler represents the sarama consumer group
type GroupHandler struct{}

// Setup is run before consumer start consuming, is normally used to setup things such as database connections
func (GroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (GroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages(), here is supposed to be what you want to
// do with the message. In this example the message will be logged with the topic name, partition and message value.
func (h GroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message topic:%q partition:%d offset:%d message: %v\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}
