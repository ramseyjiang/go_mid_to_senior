package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
	doBidirectionalStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetUnaryRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ramsey",
			LastName:  "Jiang",
		},
	}

	res, err := c.GreetUnary(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Println("Response from GreetUnary: ", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("")
	log.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetServerStreamingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ramsey",
			LastName:  "Jiang",
		},
	}

	resStream, err := c.GreetServerStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling doServerStreaming RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Println("Response from GreetServerStreaming: ", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("")
	log.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.GreetClientStreamingRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ramsey",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mamba",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Curry",
			},
		},
	}

	stream, err := c.GreetClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetClientStreaming: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending req: %v\n", req)
		err = stream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from GreetClientStreaming: %v", err)
	}

	log.Println("GreetClientStreaming Response: ", res)
}

func doBidirectionalStreaming(c greetpb.GreetServiceClient) {
	log.Println("")
	log.Println("Starting to do a Bidirectional Streaming RPC...")

	requests := []*greetpb.GreetBidirectionalStreamingRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ramsey",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mamba",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Curry",
			},
		},
	}

	stream, err := c.GreetBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetBidirectionalStreaming: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending req: %v\n", req)
		err = stream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(1000 * time.Millisecond)

		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("error while receiving response from GreetBidirectionalStreaming: %v", err)
		}

		log.Println("GreetBidirectionalStreaming Response: ", res)
	}
}
