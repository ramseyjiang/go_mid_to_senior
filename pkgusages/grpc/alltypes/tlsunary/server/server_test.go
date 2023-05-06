package main

import (
	"context"
	"errors"
	"log"
	"testing"

	tlsunary "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/tlsunary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (client tlsunary.PhoneClient, closer func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	optServer := genServerTLSOptsTest()
	baseServer := grpc.NewServer(optServer...)
	tlsunary.RegisterPhoneServer(baseServer, &phoneServer{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	optClient := genClientTLSOptsTest()
	cc, err := grpc.Dial("localhost:50055", optClient...)
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

	client = tlsunary.NewPhoneClient(cc)

	return
}

func genServerTLSOptsTest() []grpc.ServerOption {
	var opts []grpc.ServerOption
	certFile := "../cert/server.crt"
	keyFile := "../cert/server.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed loading certficates: %v\n", err)
	}
	opts = append(opts, grpc.Creds(creds))
	return opts
}

func genClientTLSOptsTest() (opts []grpc.DialOption) {
	certFile := "../cert/ca.crt"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatalf("error while loading CA trust certificate: %v\n", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	return opts
}

func TestPhoneServerGetContactName(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *tlsunary.GetContactNameResponse
		err error
	}

	tests := map[string]struct {
		in       *tlsunary.GetContactNameRequest
		expected expectation
	}{
		"MustSuccess": {
			in: &tlsunary.GetContactNameRequest{
				Number: "220123621",
			},
			expected: expectation{
				out: &tlsunary.GetContactNameResponse{
					Firstname: "Stephen",
					Lastname:  "Curry",
				},
				err: nil,
			},
		},
		"NotFoundNumber": {
			in: &tlsunary.GetContactNameRequest{
				Number: "444444444",
			},
			expected: expectation{
				out: &tlsunary.GetContactNameResponse{},
				err: errors.New("rpc error: code = Unknown desc = no contact found"),
			},
		},
		"InvalidNumber": {
			in: &tlsunary.GetContactNameRequest{
				Number: "",
			},
			expected: expectation{
				out: &tlsunary.GetContactNameResponse{},
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
		out *tlsunary.GetContactNumResponse
		err error
	}

	tests := map[string]struct {
		in       *tlsunary.GetContactNumRequest
		expected expectation
	}{
		"MustSuccess": {
			in: &tlsunary.GetContactNumRequest{
				Firstname: "Stephen",
				Lastname:  "Curry",
			},
			expected: expectation{
				out: &tlsunary.GetContactNumResponse{
					Num:    220123621,
					Result: "Owner is Stephen Curry",
				},
				err: nil,
			},
		},
		"NotFoundNumber": {
			in: &tlsunary.GetContactNumRequest{
				Firstname: "Kevin",
				Lastname:  "Durant",
			},
			expected: expectation{
				out: &tlsunary.GetContactNumResponse{},
				err: errors.New("rpc error: code = Unknown desc = no contact found"),
			},
		},
		"InvalidNameInput": {
			in: &tlsunary.GetContactNumRequest{
				Firstname: "",
				Lastname:  "",
			},
			expected: expectation{
				out: &tlsunary.GetContactNumResponse{},
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
