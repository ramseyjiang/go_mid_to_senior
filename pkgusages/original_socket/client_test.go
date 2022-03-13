package orgsocket

import (
	"net"
	"testing"
)

func TestTriggerClient(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerClient()
		})
	}
}

func Test_receiveServerData(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiveServerData(tt.args.conn)
		})
	}
}

func Test_sendClientData(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendClientData(tt.args.conn)
		})
	}
}
