package main

import (
	"context"
	"log"
	"time"

	cs "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/clientstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50057", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	client := cs.NewPhoneClient(cc)

	log.Println("Starting to do a Client Stream Phone RPC...")

	ClientStreamNumCheck(client)
}

func ClientStreamNumCheck(client cs.PhoneClient) {
	requests := []*cs.NumCheckRequest{
		{
			Number: "30",
		},
		{
			Number: "23",
		},
		{
			Number: "11",
		},
		{
			Number: "22",
		},
		{
			Number: "3",
		},
		{
			Number: "5",
		},
		{
			Number: "0",
		},
		{
			Number: "32",
		},
	}

	respStream, err := client.NumCheck(context.Background())
	if err != nil {
		log.Fatalf("error while calling Client Streaming: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending req %v\n", req)
		err = respStream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

	res, err := respStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ClientStreamServer: %v", err)
	}
	log.Println(res.CheckResult)
}
