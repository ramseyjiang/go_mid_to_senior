package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	orderRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/order/gen"
	productRPC "github.com/ramseyjiang/go_mid_to_senior/practices/grpcgw-micros/product/gen"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Register Product Service
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := productRPC.RegisterProductServiceHandlerFromEndpoint(
		ctx, mux, "product-service:50051", opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register Order Service
	err = orderRPC.RegisterOrderServiceHandlerFromEndpoint(
		ctx, mux, "order-service:50052", opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Start http service
	log.Println("gRPC Gateway running on :50050")
	http.ListenAndServe(":50050", mux)
}
