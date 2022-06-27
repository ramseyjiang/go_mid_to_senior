package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	ss "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/serverstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (client ss.PhoneClient, closer func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	ss.RegisterPhoneServer(baseServer, &phoneServer{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer = func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client = ss.NewPhoneClient(conn)

	return client, closer
}

func TestPhoneServerListContacts(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*ss.ListContactsResponse
		err error
	}

	tests := map[string]struct {
		in       *ss.ListContactsRequest
		expected expectation
	}{
		"Must_Success": {
			in: &ss.ListContactsRequest{},
			expected: expectation{
				out: []*ss.ListContactsResponse{
					{
						Firstname: "Stephen",
						Lastname:  "Curry",
						Number:    220123621,
					},
					{
						Firstname: "Klay",
						Lastname:  "Thompson",
						Number:    220123632,
					},
					{
						Firstname: "Draymond",
						Lastname:  "Green",
						Number:    220123651,
					},
					{
						Firstname: "Andrew",
						Lastname:  "Wiggins",
						Number:    220123662,
					},
					{
						Firstname: "Jorden",
						Lastname:  "Poole",
						Number:    220123671,
					},
					{
						Firstname: "Kevon",
						Lastname:  "Looney",
						Number:    220123621,
					},
					{
						Firstname: "Otto",
						Lastname:  "Porter",
						Number:    220123232,
					},
					{
						Firstname: "Garry",
						Lastname:  "Payton",
						Number:    220123355,
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.ListContacts(ctx, tt.in)

			var outs []*ss.ListContactsResponse

			for {
				o, err := out.Recv()
				if errors.Is(err, io.EOF) {
					break
				}
				outs = append(outs, o)
			}

			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if len(outs) != len(tt.expected.out) {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
				} else {
					for i, o := range outs {
						if o.Firstname != tt.expected.out[i].Firstname ||
							o.Lastname != tt.expected.out[i].Lastname ||
							o.Number != tt.expected.out[i].Number {
							t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}
