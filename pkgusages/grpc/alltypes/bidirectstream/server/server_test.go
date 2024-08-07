package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	bds "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/bidirectstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (client bds.PhoneClient, closer func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	bds.RegisterPhoneServer(baseServer, &phoneServer{})
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

	client = bds.NewPhoneClient(conn)

	return client, closer
}

func TestPhoneServerSendMsgBytes(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*bds.SendMsgBytesResponse
		err error
	}

	tests := map[string]struct {
		in       []*bds.SendMsgBytesRequest
		expected expectation
	}{
		"Must_Success": {
			in: []*bds.SendMsgBytesRequest{
				{
					Msg: []byte("Hi!"),
				},
				{
					Msg: []byte("How are you?"),
				},
				{
					Msg: []byte("Thank you!"),
				},
				{
					Msg: []byte("end"),
				},
			},
			expected: expectation{
				out: []*bds.SendMsgBytesResponse{
					{
						Msg: []byte("Hello!"),
					},
					{
						Msg: []byte("Good, good, how are you?"),
					},
					{
						Msg: []byte("Sorry, I don't understand!"),
					},
					{
						Msg: []byte("Have a good one!"),
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			outClient, err := client.SendMsgBytes(ctx)

			for _, v := range tt.in {
				if err := outClient.Send(v); err != nil {
					t.Errorf("Err -> %q", err)
				}
			}

			if err := outClient.CloseSend(); err != nil {
				t.Errorf("Err -> %q", err)
			}

			var outs []*bds.SendMsgBytesResponse
			for {
				o, err := outClient.Recv()
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
						if !bytes.Equal(o.Msg, tt.expected.out[i].Msg) {
							t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}

func TestPhoneServerSendMsgStr(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*bds.SendMsgStrResponse
		err error
	}

	tests := map[string]struct {
		in       []*bds.SendMsgStrRequest
		expected expectation
	}{
		"Must_Success": {
			in: []*bds.SendMsgStrRequest{
				{
					Msg: "Hey Green. What's up?",
				},
				{
					Msg: "Not bad, bro. Do you want to have a new contract?",
				},
				{
					Msg: "Please be serious. How much do you want?",
				},
				{
					Msg: "That will be great for you. But we don't have enough space if you get that.",
				},
				{
					Msg: "Ok, let's hold to see what will happen. Have a good one!",
				},
			},
			expected: expectation{
				out: []*bds.SendMsgStrResponse{
					{
						Msg: "Great, how are you, Mills!",
					},
					{
						Msg: "Of course, bro! The largest I can get, please.",
					},
					{
						Msg: "The largest I can get, please.",
					},
					{
						Msg: "Sorry, I don't understand! I think you can.",
					},
					{
						Msg: "Have a good one!",
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			outClient, err := client.SendMsgStr(ctx)

			for _, v := range tt.in {
				if err := outClient.Send(v); err != nil {
					t.Errorf("Err -> %q", err)
				}
			}

			if err := outClient.CloseSend(); err != nil {
				t.Errorf("Err -> %q", err)
			}

			var outs []*bds.SendMsgStrResponse
			for {
				o, err := outClient.Recv()
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
						if o.Msg != tt.expected.out[i].Msg {
							t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}
