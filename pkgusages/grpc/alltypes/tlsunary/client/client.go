package main

import (
	"context"
	"log"
	"strconv"

	"google.golang.org/grpc/credentials"

	tlsunary "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/tlsunary/proto"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	certFile := "cert/ca.crt"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")

	if err != nil {
		log.Fatalf("error while loading CA trust certificate: %v\n", err)
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	cc, err := grpc.Dial("localhost:50055", opts...)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	client := tlsunary.NewPhoneClient(cc)

	log.Println("Starting to do a Unary Phone RPC...")

	GetServerContactName(client)
	GetServerContactNum(client)
	ListContacts(client)
}

func GetServerContactName(client tlsunary.PhoneClient) {
	req := &tlsunary.GetContactNameRequest{
		Number: "220123651",
	}

	res, err := client.GetContactName(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}

	log.Println("Response from GetContactName: ", res.Firstname+" "+res.Lastname)
}

func GetServerContactNum(client tlsunary.PhoneClient) {
	req := &tlsunary.GetContactNumRequest{
		Firstname: "Klay",
		Lastname:  "Thompson",
	}

	res, err := client.GetContactNum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}

	log.Println("Response from GetContactNum: ", strconv.FormatUint(res.Num, 10)+" "+res.Result)
}

func ListContacts(client tlsunary.PhoneClient) {
	res, err := client.ListContacts(context.Background(), &tlsunary.ListContactsRequest{})
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}
	log.Println("Response from ListContacts: sum is ", res.Sum)
	log.Println("contacts list is ", res.Contacts)
}
