package main

import (
	"context"
	"log"
	"net"
	"time"

	orderRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/order/gen"
	"google.golang.org/grpc"
)

type server struct {
	orderRPC.UnimplementedOrderServiceServer
	// Add connect db
}

func (s *server) CreateOrder(ctx context.Context, req *orderRPC.CreateOrderRequest) (*orderRPC.CreateOrderResponse, error) {
	return &orderRPC.CreateOrderResponse{
		OrderId: "order-" + time.Now().Format("20060102150405"),
	}, nil
}

func (s *server) GetOrders(ctx context.Context, req *orderRPC.GetOrdersRequest) (*orderRPC.GetOrdersResponse, error) {
	orders := []*orderRPC.OrderDetail{
		{
			OrderId:   "order-123",
			UserId:    req.UserId,
			Status:    "created",
			CreatedAt: time.Now().Unix(),
		},
		{
			OrderId:   "order-456",
			UserId:    req.UserId,
			Status:    "paid",
			CreatedAt: time.Now().Add(-24 * time.Hour).Unix(),
		},
	}

	return &orderRPC.GetOrdersResponse{Orders: orders}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50052")
	s := grpc.NewServer()
	orderRPC.RegisterOrderServiceServer(s, &server{})
	log.Println("Order Service run on :50052")
	s.Serve(lis)
}
