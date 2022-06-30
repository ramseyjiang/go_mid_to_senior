package main

import (
	"errors"
	"io"
	"log"
	"net"

	bds "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/bidirectstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	bds.PhoneServer
	calls []string
}

func (p *phoneServer) SendMessage(stream bds.Phone_SendMessageServer) (err error) {
	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		p.calls = append(p.calls, string(req.Msg))
		if p.calls[len(p.calls)-1] == "." {
			for _, m := range p.calls {
				// time.Sleep(time.Second)
				switch m {
				case ".":
				case "Hi!":
					_ = stream.Send(&bds.SendMessageResponse{
						Msg: []byte("Hello!"),
					})
				case "How are you?":
					_ = stream.Send(&bds.SendMessageResponse{
						Msg: []byte("Fine, you?"),
					})
				case "See you later":
					_ = stream.Send(&bds.SendMessageResponse{
						Msg: []byte("See you!"),
					})
				default:
					_ = stream.Send(&bds.SendMessageResponse{
						Msg: []byte("Sorry, I don't understand :/"),
					})
				}
			}
			p.calls = p.calls[:0]
		}
	}
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
