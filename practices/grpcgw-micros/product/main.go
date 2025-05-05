package main

import (
	"context"
	"log"
	"net"

	productRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/product/gen"
	"google.golang.org/grpc"
)

type server struct {
	productRPC.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *productRPC.ProductRequest) (*productRPC.ProductResponse, error) {
	return &productRPC.ProductResponse{
		Id:    "p-123",
		Name:  "Demo Product",
		Price: 99.99,
	}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	productRPC.RegisterProductServiceServer(s, &server{})
	log.Println("Product gRPC Server running on :50051")
	s.Serve(listen)
}
