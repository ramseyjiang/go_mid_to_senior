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

	BidirectionalStreamSendBytesMsg(client)
	BidirectionalStreamSendStrMsg(client)
}

func BidirectionalStreamSendBytesMsg(client bds.PhoneClient) {
	requests := []*bds.SendMsgBytesRequest{
		{
			Msg: []byte("Hi!"),
		},
		{
			Msg: []byte("How are you?"),
		},
		{
			Msg: []byte("end"),
		},
		{
			Msg: []byte("Thank you!"),
		},
		{
			Msg: []byte("end"),
		},
	}

	respStream, err := client.SendMsgBytes(context.Background())
	if err != nil {
		log.Fatalf("error while calling Bidirectional Streaming: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending bytes req: %v\n", req)
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

	var responses []*bds.SendMsgBytesResponse
	for {
		resp, err := respStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		responses = append(responses, resp)
		time.Sleep(1 * time.Second)
		log.Println("SendMessageResponse:", resp)
	}
	log.Println("All SendMessageResponse", responses)
}

func BidirectionalStreamSendStrMsg(client bds.PhoneClient) {
	requests := []*bds.SendMsgStrRequest{
		{
			Msg: "Hi!",
		},
		{
			Msg: "How are you?",
		},
		{
			Msg: "end",
		},
		{
			Msg: "Thank you!",
		},
		{
			Msg: "end",
		},
	}

	respStream, err := client.SendMsgStr(context.Background())
	if err != nil {
		log.Fatalf("error while calling Bidirectional Streaming: %v", err)
	}

	var responses []*bds.SendMsgStrResponse
	for _, req := range requests {
		log.Printf("Sending str req: %v\n", req)
		err = respStream.Send(req)
		time.Sleep(1 * time.Second)
		if err != nil {
			return
		}
		resp, err := respStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		time.Sleep(1 * time.Second)
		log.Println("SendMessageResponse:", resp)
		time.Sleep(1 * time.Second)
		responses = append(responses, resp)
	}
	log.Println("All SendMessageResponse", responses)

	err = respStream.CloseSend()
	if err != nil {
		log.Fatalf("error while receiving response from Bidirectional Streaming: %v", err)
	}
}
