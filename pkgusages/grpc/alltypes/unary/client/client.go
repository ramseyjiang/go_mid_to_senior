package main

import (
	"context"
	"log"
	"strconv"

	unary "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	client := unary.NewPhoneClient(cc)

	log.Println("Starting to do a Unary Phone RPC...")

	GetServerContactName(client)
	GetServerContactNum(client)
}

func GetServerContactName(client unary.PhoneClient) {
	req := &unary.GetContactNameRequest{
		Number: "220123651",
	}

	res, err := client.GetContactName(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}

	log.Println("Response from GetContactName: ", res.Firstname+" "+res.Lastname)
}

func GetServerContactNum(client unary.PhoneClient) {
	req := &unary.GetContactNumRequest{
		Firstname: "Klay",
		Lastname:  "Thompson",
	}

	res, err := client.GetContactNum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}

	log.Println("Response from GetContactNum: ", strconv.FormatUint(res.Num, 10)+" "+res.Result)
}
