package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "myTopic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:19091", topic, partition)

	if err != nil {
		log.Println("failed to dial leader:", err)
	}

	if err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatalf("fatal error is:%v\n", err)
	}
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		log.Println("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Println("failed to close writer:", err)
	}

	log.Println("Conn Close")
	log.Println(err)
}
