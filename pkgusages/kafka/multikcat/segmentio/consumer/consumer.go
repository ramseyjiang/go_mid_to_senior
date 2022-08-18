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

	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Println("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Println("failed to close connection:", err)
	}
}
