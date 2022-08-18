package consumer

import (
	"log"

	"github.com/Shopify/sarama"
)

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
