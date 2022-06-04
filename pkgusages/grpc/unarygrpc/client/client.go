package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	greetpb2 "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unarygrpc/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err1 := cc.Close()
		if err1 != nil {
			fmt.Println(err1)
		}
	}(cc)

	c := greetpb2.NewGreetServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
}

func doClientStreaming(c greetpb2.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb2.LongGreetingRequest{
		{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ramsey",
			},
		},
		{
			Greeting: &greetpb2.Greeting{
				FirstName: "Mamba",
			},
		},
		{
			Greeting: &greetpb2.Greeting{
				FirstName: "Mamba",
			},
		},
		{
			Greeting: &greetpb2.Greeting{
				FirstName: "Mamba",
			},
		},
	}

	stream, err := c.LongGreeting(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		err1 := stream.Send(req)
		if err1 != nil {
			return
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet Response: %v\n", res)
}

func doServerStreaming(c greetpb2.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb2.GreetManyTimesRequest{
		Greeting: &greetpb2.Greeting{
			FirstName: "Ramsey",
			LastName:  "Jiang",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
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

		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}

func doUnary(c greetpb2.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb2.GreetRequest{
		Greeting: &greetpb2.Greeting{
			FirstName: "Ramsey",
			LastName:  "Jiang",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	fmt.Printf("Response from Greet: %v", res.Result)
}
