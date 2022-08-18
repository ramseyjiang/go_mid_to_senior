package producer

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers = []string{"localhost:9093"}
	KafkaTopic   = "sarama_topic"
	enqueued     int
)

func StartProducer() {
	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	produceMessages(producer, signals)

	log.Printf("Kafka AsyncProducer finished with %d messages produced.", enqueued)
}

// setupProducer will create a AsyncProducer and returns it
func setupProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	return sarama.NewAsyncProducer(kafkaBrokers, config)
}

// produceMessages will send 'testing 123' to KafkaTopic each second, until receive an os signal to stop e.g. control + c
// by the user in terminal
func produceMessages(producer sarama.AsyncProducer, signals chan os.Signal) {
	for {
		time.Sleep(time.Second)
		message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder("testing 1234")}
		select {
		case producer.Input() <- message:
			enqueued++
			log.Println("New Message produced")
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			return
		}
	}
}
