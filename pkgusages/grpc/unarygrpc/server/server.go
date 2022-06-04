package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	greetpb2 "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unarygrpc/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb2.GreetRequest) (*greetpb2.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v %v,\n", req, ctx)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	res := &greetpb2.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb2.GreetManyTimesRequest, stream greetpb2.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 3; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb2.GreetManyTimesResponse{
			Result: result,
		}
		err1 := stream.Send(res)
		if err1 != nil {
			return err1
		}
		log.Printf("Sent: %v", res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreeting(stream greetpb2.GreetService_LongGreetingServer) error {
	fmt.Printf("LongGreeting function was invoked with a streaming request\n")

	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb2.LongGreetingResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Print("Server started")
	s := grpc.NewServer()
	greetpb2.RegisterGreetServiceServer(s, &server{})

	if err1 := s.Serve(listener); err1 != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
