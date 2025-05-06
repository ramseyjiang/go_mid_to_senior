package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	productRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/product/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productRPC.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *productRPC.ProductRequest) (*productRPC.ProductResponse, error) {
	// Create a 2 seconds timeout using context
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-timeoutCtx.Done():
		// No timeout logic
		return &productRPC.ProductResponse{
			ID:    "p-1234",
			Name:  "Demo Product",
			Price: 99.99,
		}, nil
	case <-time.After(3 * time.Second): // Over 3 seconds trigger the timeout
		// Trigger by timeout or cancel
		if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "Request process timeout.")
		}
		return nil, status.Error(codes.Canceled, "Request Cancel")
	}
}

func main() {
	listen, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	productRPC.RegisterProductServiceServer(s, &server{})
	log.Println("Product gRPC Server running on :50051")
	s.Serve(listen)
}
