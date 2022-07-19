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

	ClientStreamCallRecord(client)
}

func ClientStreamCallRecord(client cs.PhoneClient) {
	requests := []*cs.CallRecordRequest{
		{
			Number: "11111111111",
		},
		{
			Number: "22222222222",
		},
		{
			Number: "33333333333",
		},
	}

	respStream, err := client.CallRecord(context.Background())
	if err != nil {
		log.Fatalf("error while calling Client Streaming: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending req: %v\n", req)
		err = respStream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(5 * time.Second)
	}

	res, err := respStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from GreetClientStreaming: %v", err)
	}

	log.Println("RecordCallHistoryResponse: CallCount is", res.CallCount)
	log.Println("RecordCallHistoryResponse: DurationCall time is", res.DurationCall)
}
