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
	ServerStreamPageContacts(client)
}

func ServerStreamListContacts(client ss.PhoneClient) {
	respStream, err := client.AllContacts(context.Background(), &ss.AllContactsRequest{})
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

		log.Println("Response from Server Stream Phone RPC: number is", msg.Contact.Number,
			"firstname is ", msg.Contact.Firstname, "lastname is", msg.Contact.Lastname)
	}
}

func ServerStreamPageContacts(client ss.PhoneClient) {
	respStream, err := client.PageContacts(context.Background(), &ss.PageContactsRequest{PageSize: "3"})
	if err != nil {
		log.Fatalf("error while calling Server Stream Phone RPC: %v", err)
	}

	for {
		resp, err := respStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Println("Response from Server Stream Phone RPC: CurrentPage is", resp.CurrentPage,
			"PageSize is", resp.PageSize,
			"PageAmount is", resp.PageAmount,
			"Contact[0] Firstname is", resp.Contact[0].Firstname)
	}
}
