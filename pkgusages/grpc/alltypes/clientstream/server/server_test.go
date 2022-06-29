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
func TestPhoneServerRecordCallHistory(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *cs.RecordCallHistoryResponse
		err error
	}

	tests := map[string]struct {
		in       []*cs.RecordCallHistoryRequest
		expected expectation
	}{
		"Must_Success": {
			in: []*cs.RecordCallHistoryRequest{
				{
					Number: "11111111111",
				},
				{
					Number: "22222222222",
				},
				{
					Number: "33333333333",
				},
			},
			expected: expectation{
				out: &cs.RecordCallHistoryResponse{
					CallCount: 3,
				},
				err: nil,
			},
		},
		"Empty_Request": {
			in: []*cs.RecordCallHistoryRequest{},
			expected: expectation{
				out: &cs.RecordCallHistoryResponse{
					CallCount: 0,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			outClient, err := client.RecordCallHistory(ctx)

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
				if tt.expected.out.CallCount != out.CallCount {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

			if err := outClient.CloseSend(); err != nil {
				t.Errorf("Err -> %q", err)
			}
		})
	}
}
