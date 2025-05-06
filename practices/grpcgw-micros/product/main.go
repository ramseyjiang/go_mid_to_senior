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
	// Set 2 seconds timeout using context
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Use a channel to receive the result of the actual product retrieval.
	productChan := make(chan *productRPC.ProductResponse, 1) // Buffered channel to avoid goroutine leak
	errChan := make(chan error, 1)

	// Perform the actual product retrieval in a separate goroutine.
	go func() {
		// Simulate fetching product data (TODO: replace with your actual database/service call).
		//  If you are using a database driver, make sure the database query uses the timeoutContext.
		//  If you are calling another service, make sure the http/grpc client uses the timeoutContext.
		time.Sleep(1 * time.Second) // Simulate a 1-second delay.  Replace this with your actual data fetching logic.
		product := &productRPC.ProductResponse{
			ProductID: int32(1234),
			Name:      "Demo Product",
			Price:     19.99,
		}
		productChan <- product
		errChan <- nil // Signal success
	}()

	select {
	case <-timeoutCtx.Done():
		// Trigger by timeout or cancel
		if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "Request process timeout.")
		}
		return nil, status.Error(codes.Canceled, "Request Cancel")
	case product := <-productChan: // Receive the product
		err := <-errChan
		if err != nil {
			return nil, err
		}
		return product, nil
	}
}

func main() {
	listen, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	productRPC.RegisterProductServiceServer(s, &server{})
	log.Println("Product gRPC Server running on :50051")
	s.Serve(listen)
}
