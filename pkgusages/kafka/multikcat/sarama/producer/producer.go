package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers = []string{"0.0.0.0:19091"}
	KafkaTopic   = "myTopic"
	enqueued     int
)

func main() {
	producer, err := sarama.NewAsyncProducer(kafkaBrokers, sarama.NewConfig())
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

// produceMessages will send 'testing 1234' to KafkaTopic each second, until receive an os signal to stop e.g. control + c
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
