package main

import (
	"context"
	"log"
	"time"

	cal "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unary/calnogateway/proto"
	"google.golang.org/grpc"
)

func main() {
	// create client connection
	conn, err := grpc.Dial("0.0.0.0:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := cal.NewCalculateServiceClient(conn)

	// Contact the server and print out its response.
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	doAdd(client)
	doSub(client)
	doMul(client)
}

func doAdd(client cal.CalculateServiceClient) {
	addResp, err := client.Add(context.Background(), &cal.Request{A: 1, B: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Add result is ", addResp.Result)
}

func doSub(client cal.CalculateServiceClient) {
	subResp, err := client.Add(context.Background(), &cal.Request{A: 5, B: 3})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sub result is ", subResp.Result)
}

func doMul(client cal.CalculateServiceClient) {
	mulResp, err := client.Mul(context.Background(), &cal.Request{A: 5, B: 3})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mul result is ", mulResp.Result)
}
