package main

import (
	"context"
	"log"
	"net"

	cal "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unary/calnogateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CalService struct {
	cal.CalculateServiceServer
}

func (s *CalService) Add(ctx context.Context, request *cal.Request) (*cal.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b
	return &cal.Response{Result: result}, nil
}

func (s *CalService) Sub(ctx context.Context, request *cal.Request) (*cal.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b
	return &cal.Response{Result: result}, nil
}

func (s *CalService) Mul(ctx context.Context, request *cal.Request) (*cal.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b
	return &cal.Response{Result: result}, nil
}

func main() {
	// listen port
	lis, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("list port err: %v", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer()

	// register service into grpc server
	cal.RegisterCalculateServiceServer(grpcServer, &CalService{})
	reflection.Register(grpcServer)

	log.Printf("listening at %v", lis.Addr())

	// listen port
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc serve err: %v", err)
	}
}
