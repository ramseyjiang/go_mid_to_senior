package main

import (
	"context"
	"net"

	proto2 "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unary/calwithgateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto2.RegisterCalculateServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto2.Request) (*proto2.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto2.Response{Result: result}, nil
}

func (s *server) Sub(ctx context.Context, request *proto2.Request) (*proto2.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b

	return &proto2.Response{Result: result}, nil
}

func (s *server) Mul(ctx context.Context, request *proto2.Request) (*proto2.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto2.Response{Result: result}, nil
}
