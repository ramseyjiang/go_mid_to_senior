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

// use slice and for loop to simulate server-side streaming.
// In the for loop, it will get replies using out.Recv() function and store them at outs slice.
func TestPhoneServerAllContacts(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*ss.AllContactsResponse
		err error
	}

	tests := map[string]struct {
		in       *ss.AllContactsRequest
		expected expectation
	}{
		"Must_Success": {
			in: &ss.AllContactsRequest{},
			expected: expectation{
				out: []*ss.AllContactsResponse{
					{
						Contact: &ss.Contact{
							Firstname: "Stephen",
							Lastname:  "Curry",
							Number:    220123621,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Klay",
							Lastname:  "Thompson",
							Number:    220123632,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Draymond",
							Lastname:  "Green",
							Number:    220123651,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Andrew",
							Lastname:  "Wiggins",
							Number:    220123662,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Jorden",
							Lastname:  "Poole",
							Number:    220123671,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Kevon",
							Lastname:  "Looney",
							Number:    220123621,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Otto",
							Lastname:  "Porter",
							Number:    220123232,
						},
					},
					{
						Contact: &ss.Contact{
							Firstname: "Garry",
							Lastname:  "Payton",
							Number:    220123355,
						},
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.AllContacts(ctx, tt.in)

			var outs []*ss.AllContactsResponse

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
						if o.Contact.Firstname != tt.expected.out[i].Contact.Firstname ||
							o.Contact.Lastname != tt.expected.out[i].Contact.Lastname ||
							o.Contact.Number != tt.expected.out[i].Contact.Number {
							t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}

func TestPhoneServerPageContacts(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*ss.PageContactsResponse
		err error
	}

	tests := map[string]struct {
		in       *ss.PageContactsRequest
		expected expectation
	}{
		"Must_Success": {
			in: &ss.PageContactsRequest{
				PageSize: "4",
			},
			expected: expectation{
				out: []*ss.PageContactsResponse{
					{
						PageAmount:  2,
						PageSize:    4,
						CurrentPage: 1,
						Contact: []*ss.Contact{
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
						},
					},
					{
						PageAmount:  2,
						PageSize:    4,
						CurrentPage: 2,
						Contact: []*ss.Contact{
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
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.PageContacts(ctx, tt.in)

			var outs []*ss.PageContactsResponse

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
						if o.CurrentPage != tt.expected.out[i].CurrentPage ||
							o.PageSize != tt.expected.out[i].PageSize ||
							o.PageAmount != tt.expected.out[i].PageAmount {
							t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}
