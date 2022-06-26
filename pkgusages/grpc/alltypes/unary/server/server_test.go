package main

import (
	"context"
	"errors"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	unary "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/unary/proto"
)

func server(ctx context.Context) (client unary.PhoneClient, closer func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	unary.RegisterPhoneServer(baseServer, &phoneServer{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	cc, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer = func() {
		err = lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client = unary.NewPhoneClient(cc)

	return
}

func TestPhoneServerGetContactName(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *unary.GetContactNameResponse
		err error
	}

	tests := map[string]struct {
		in       *unary.GetContactNameRequest
		expected expectation
	}{
		"MustSuccess": {
			in: &unary.GetContactNameRequest{
				Number: "220123621",
			},
			expected: expectation{
				out: &unary.GetContactNameResponse{
					Firstname: "Stephen",
					Lastname:  "Curry",
				},
				err: nil,
			},
		},
		"NotFoundNumber": {
			in: &unary.GetContactNameRequest{
				Number: "444444444",
			},
			expected: expectation{
				out: &unary.GetContactNameResponse{},
				err: errors.New("rpc error: code = Unknown desc = no contact found"),
			},
		},
		"InvalidNumber": {
			in: &unary.GetContactNameRequest{
				Number: "",
			},
			expected: expectation{
				out: &unary.GetContactNameResponse{},
				err: errors.New("rpc error: code = Unknown desc = invalid number"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetContactName(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Firstname != out.Firstname || tt.expected.out.Lastname != out.Lastname {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}
		})
	}
}

func TestPhoneServerGetContactNum(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *unary.GetContactNumResponse
		err error
	}

	tests := map[string]struct {
		in       *unary.GetContactNumRequest
		expected expectation
	}{
		"MustSuccess": {
			in: &unary.GetContactNumRequest{
				Firstname: "Stephen",
				Lastname:  "Curry",
			},
			expected: expectation{
				out: &unary.GetContactNumResponse{
					Num:    220123621,
					Result: "Owner is Stephen Curry",
				},
				err: nil,
			},
		},
		"NotFoundNumber": {
			in: &unary.GetContactNumRequest{
				Firstname: "Kevin",
				Lastname:  "Durant",
			},
			expected: expectation{
				out: &unary.GetContactNumResponse{},
				err: errors.New("rpc error: code = Unknown desc = no contact found"),
			},
		},
		"InvalidNameInput": {
			in: &unary.GetContactNumRequest{
				Firstname: "",
				Lastname:  "",
			},
			expected: expectation{
				out: &unary.GetContactNumResponse{},
				err: errors.New("rpc error: code = Unknown desc = invalid firstname and lastname input"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetContactNum(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Num != out.Num || tt.expected.out.Result != out.Result {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}
		})
	}
}
