package main

import (
	"flag"
	"fmt"

	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/kafka/sarama/consumer"
	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/kafka/sarama/producer"
)

func main() {
	// Flags
	consumerFlagValue := flag.Bool("consumer", false, "    Use this flag to start a Kafka Consumer")
	producerFlagValue := flag.Bool("producer", false, "    Use this flag to start a Kafka Producer")
	stringFlagValue := flag.String("all", "", "    Use this flag with either \"consumer\" or \"producer\"")

	// Flag Processing
	flag.Parse()

	// Decision Time
	if *producerFlagValue == true {
		producer.StartProducer()
	} else if *consumerFlagValue == true {
		consumer.StartConsumer()
	} else if *stringFlagValue == "consumer" {
		consumer.StartConsumer()
	} else if *stringFlagValue == "producer" {
		producer.StartProducer()
	} else {
		fmt.Print("Usage: \n -c     Use this flag to start a Kafka Consumer\n -p     Use this flag to start a Kafka Producer\n -a     Use this flag with either \"consumer\" or \"producer\"\n")
	}
}
