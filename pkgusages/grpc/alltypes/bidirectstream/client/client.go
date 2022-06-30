package main

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	bds "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/bidirectstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50058", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	client := bds.NewPhoneClient(cc)

	log.Println("Starting to do a Bidirectional Stream Phone RPC...")

	BidirectionalStreamSendMsg(client)
}

func BidirectionalStreamSendMsg(client bds.PhoneClient) {
	requests := []*bds.SendMessageRequest{
		{
			Msg: []byte("Hi!"),
		},
		{
			Msg: []byte("How are you?"),
		},
		{
			Msg: []byte("Thank you!"),
		},
		{
			Msg: []byte("."),
		},
	}

	respStream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("error while calling Bidirectional Streaming: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending req: %v\n", req)
		err = respStream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

	err = respStream.CloseSend()
	if err != nil {
		log.Fatalf("error while receiving response from Bidirectional Streaming: %v", err)
	}

	var responses []*bds.SendMessageResponse
	for {
		o, err := respStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		responses = append(responses, o)
	}

	log.Println("All SendMessageResponse", responses)
	for key, val := range responses {
		log.Println("SendMessageResponse: key is", key)
		log.Println("SendMessageResponse: val is", val)
	}
}
