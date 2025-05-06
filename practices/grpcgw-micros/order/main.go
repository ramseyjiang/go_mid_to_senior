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
		OrderID: "order-" + time.Now().Format("20060102150405"),
	}, nil
}

func (s *server) GetOrders(ctx context.Context, req *orderRPC.GetOrdersRequest) (*orderRPC.GetOrdersResponse, error) {
	orders := []*orderRPC.OrderDetail{
		{
			OrderID:   "order-123",
			UserID:    req.UserID,
			Status:    orderRPC.OrderStatus_ORDER_STATUS_CREATED,
			CreatedAt: time.Now().Unix(),
		},
		{
			OrderID:   "order-456",
			UserID:    req.UserID,
			Status:    orderRPC.OrderStatus_ORDER_STATUS_PAID,
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
