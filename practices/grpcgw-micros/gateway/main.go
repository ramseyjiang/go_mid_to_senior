package main

import (
	"context"
	"log"
	"net/http"

	orderRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/order/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	productRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/product/gen"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	// 注册 Product 服务端点
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := productRPC.RegisterProductServiceHandlerFromEndpoint(
		ctx, mux, "product-service:50051", opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 在网关中注册 Order 服务
	err = orderRPC.RegisterOrderServiceHandlerFromEndpoint(
		ctx, mux, "order-service:50052", opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 启动 HTTP 服务
	log.Println("gRPC Gateway running on :50050")
	http.ListenAndServe(":50050", mux)
}
