package main

import (
	"errors"
	"io"
	"log"
	"net"
	"time"

	bds "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/bidirectstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	bds.PhoneServer
	calls []string
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50058")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")
	s := grpc.NewServer()
	bds.RegisterPhoneServer(s, &phoneServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (p *phoneServer) SendMsgBytes(stream bds.Phone_SendMsgBytesServer) (err error) {
	log.Println("Server SendMsgBytes started")
	for {
		req, err := stream.Recv()
		if err != nil {
			// io.EOF is used to check the end of the stream.
			// Checking for the client stream message in the infinite loop and sending the response messages,
			// once the client ends the stream then we break the loop.
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		p.calls = append(p.calls, string(req.Msg))
		log.Println("record calls received number is ", len(p.calls))

		if p.calls[len(p.calls)-1] == "end" {
			for _, msg := range p.calls {
				// time.Sleep(time.Second)
				switch msg {
				case "end":
					_ = stream.Send(&bds.SendMsgBytesResponse{
						Msg: []byte("Have a good one!"),
					})
				case "Hi!":
					_ = stream.Send(&bds.SendMsgBytesResponse{
						Msg: []byte("Hello!"),
					})
				case "How are you?":
					_ = stream.Send(&bds.SendMsgBytesResponse{
						Msg: []byte("Good, good, how are you?"),
					})
				case "See you later":
					_ = stream.Send(&bds.SendMsgBytesResponse{
						Msg: []byte("See you!"),
					})
				default:
					_ = stream.Send(&bds.SendMsgBytesResponse{
						Msg: []byte("Sorry, I don't understand :/"),
					})
				}
			}
			p.calls = p.calls[:0]
		}
	}
}

func (p *phoneServer) SendMsgStr(stream bds.Phone_SendMsgStrServer) (err error) {
	log.Println("Server SendMsgStr started")
	for {
		req, err := stream.Recv()
		if err != nil {
			// io.EOF is used to check the end of the stream.
			// Checking for the client stream message in the infinite loop and sending the response messages,
			// once the client ends the stream then we break the loop.
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		p.calls = append(p.calls, req.Msg)
		log.Println("record calls received number is ", len(p.calls))

		for _, msg := range p.calls {
			time.Sleep(time.Second)
			switch msg {
			case "end":
				_ = stream.Send(&bds.SendMsgStrResponse{
					Msg: "Have a good one!",
				})
			case "Hi!":
				_ = stream.Send(&bds.SendMsgStrResponse{
					Msg: "Hello!",
				})
			case "How are you?":
				_ = stream.Send(&bds.SendMsgStrResponse{
					Msg: "Good, good, how are you?",
				})
			case "See you later":
				_ = stream.Send(&bds.SendMsgStrResponse{
					Msg: "See you!",
				})
			default:
				_ = stream.Send(&bds.SendMsgStrResponse{
					Msg: "Sorry, I don't understand :/",
				})
			}
		}
		p.calls = p.calls[:0]
	}
}
