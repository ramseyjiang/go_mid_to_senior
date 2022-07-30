package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	cs "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/clientstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (client cs.PhoneClient, closer func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	cs.RegisterPhoneServer(baseServer, &phoneServer{})
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

	client = cs.NewPhoneClient(conn)

	return client, closer
}

// use for loop to simulate client-side streaming.
// In the for loop, it will send requests to the RecordCallHistory client using outClient.Send().
// Once all requests have been sent, we will close the sending.
func TestPhoneServerNumCheck(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*cs.NumCheckResponse
		err error
	}

	tests := map[string]struct {
		in       []*cs.NumCheckRequest
		expected expectation
	}{
		"Must_Success": {
			in: []*cs.NumCheckRequest{
				{
					Number: "30",
				},
				{
					Number: "11",
				},
				{
					Number: "0",
				},
			},
			expected: expectation{
				out: []*cs.NumCheckResponse{
					{
						CheckResult: []*cs.Result{
							{
								Msg:    "Stephen Curry phone is 220123621.",
								Status: "Stay",
							},
							{
								Msg:    "Klay Thompson phone is 220123632.",
								Status: "Stay",
							},
							{
								Msg:    "0 Has joined another team.",
								Status: "Left",
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
			outClient, err := client.NumCheck(ctx)

			for _, v := range tt.in {
				if err := outClient.Send(v); err != nil {
					t.Errorf("Err -> %q", err)
				}
			}

			out, err := outClient.CloseAndRecv()
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				for i, o := range out.CheckResult {
					if o.Status != tt.expected.out[0].CheckResult[i].Status &&
						o.Msg != tt.expected.out[0].CheckResult[i].Msg {
						t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
					}
				}
			}

			if err := outClient.CloseSend(); err != nil {
				t.Errorf("Err -> %q", err)
			}
		})
	}
}
