package main

import (
	"errors"
	"io"
	"log"
	"net"
	"time"

	cs "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/clientstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	cs.PhoneServer
}

func (p *phoneServer) RecordCallHistory(stream cs.Phone_RecordCallHistoryServer) (err error) {
	var callCount int32

	start := time.Now()

	for {
		_, err = stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return stream.SendAndClose(&cs.RecordCallHistoryResponse{
					CallCount:    callCount,
					DurationCall: int32(time.Since(start)),
				})
			}
			return err
		}
		callCount++
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50057")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")
	s := grpc.NewServer()
	cs.RegisterPhoneServer(s, &phoneServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
