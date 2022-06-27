package main

import (
	"context"
	"io"
	"log"

	ss "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/serverstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	client := ss.NewPhoneClient(cc)

	log.Println("Starting to do a Server Stream Phone RPC...")

	ServerStreamListContacts(client)
}

func ServerStreamListContacts(client ss.PhoneClient) {
	respStream, err := client.ListContacts(context.Background(), &ss.ListContactsRequest{})
	if err != nil {
		log.Fatalf("error while calling Server Stream Phone RPC: %v", err)
	}

	for {
		msg, err := respStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Println("Response from Server Stream Phone RPC: number is", msg.Number, "firstname is ", msg.Firstname, "lastname is", msg.Lastname)
	}
}
