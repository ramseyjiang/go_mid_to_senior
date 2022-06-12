package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) GreetUnary(ctx context.Context, req *proto.GreetUnaryRequest) (*proto.GreetUnaryResponse, error) {
	log.Printf("GreetUnary server was invoked with %v,\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	res := &proto.GreetUnaryResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetServerStreaming(req *proto.GreetServerStreamingRequest, stream proto.GreetService_GreetServerStreamingServer) error {
	log.Println("")
	log.Printf("GreetServerStreaming server was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 2; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &proto.GreetServerStreamingResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		log.Printf("Sent: %v", res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) GreetClientStreaming(stream proto.GreetService_GreetClientStreamingServer) error {
	log.Println("")
	log.Printf("GreetClientStreaming server was invoked with a streaming request\n")

	result := ""
	for {
		req, err := stream.Recv()

		// receiving the stream message in the infinite loop and priting to the terminal
		// once err equals io.EOF (end of the stream) we send the single response message
		// and close the stream then break the loop.
		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetClientStreamingResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result = "Hello " + firstName + "! "
		log.Println(result)
	}
	return nil
}

func (*server) GreetBidirectionalStreaming(stream proto.GreetService_GreetBidirectionalStreamingServer) error {
	log.Println("")
	log.Printf("GreetBidirectionalStreaming server was invoked with a streaming request\n")

	for {
		// Recv() — to receive the client stream messages
		req, err := stream.Recv()

		// io.EOF is used to check the end of the stream.
		// Checking for the client stream message in the infinite loop and sending the response messages,
		// once the client ends the stream then we break the loop.
		if err == io.EOF {
			break
		}

		if err != nil {
			// Because it is an infinite loop, after the last req received,
			// it will print an error info "Error in stream Recv:  rpc error: code = Canceled desc = context canceled".
			// But after the new request comes, it can keep working.
			log.Println("Error in stream Recv: ", err)
			break
		}

		log.Println("Request - ", req)
		result := "Hello " + req.GetGreeting().GetFirstName()
		res := &proto.GreetBidirectionalStreamingResponse{
			Result: result,
		}

		// Send() — to send the stream messages
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("Error in stream close: %v", err)
		}
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")
	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
