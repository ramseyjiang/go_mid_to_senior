package main

import (
	"context"
	"log"
	"net"
	"time"

	orderRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/order/gen"
	productRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/product/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type orderServer struct {
	orderRPC.UnimplementedOrderServiceServer
	productClient productRPC.ProductServiceClient // gRPC 客户端
}

// NewOrderServer should Connect the Product Service at first.
func NewOrderServer() *orderServer {
	conn, err := grpc.Dial(
		"product-service:50051", // 通过 Docker 服务名访问
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Connect Product Service Failed: %v", err)
	}
	return &orderServer{
		productClient: productRPC.NewProductServiceClient(conn),
	}
}

func (s *orderServer) CreateOrder(ctx context.Context, req *orderRPC.CreateOrderRequest) (*orderRPC.CreateOrderResponse, error) {
	// 调用 Product Service 获取商品信息（带超时控制）
	productCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	productResp, err := s.productClient.GetProduct(productCtx, &productRPC.ProductRequest{
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, err
	}

	// 生成订单逻辑（此处简化为示例）
	return &orderRPC.CreateOrderResponse{
		OrderId: "order-123",
		Product: productResp,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50052") // Order 服务端口
	s := grpc.NewServer()
	orderRPC.RegisterOrderServiceServer(s, NewOrderServer())
	log.Println("Order gRPC Server On Port :50052")
	s.Serve(lis)
}
