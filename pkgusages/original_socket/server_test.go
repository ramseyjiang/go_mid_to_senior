package orgsocket

import (
	"net"
	"testing"
)

func TestTriggerServer(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerServer()
		})
	}
}

func Test_receiveClientData(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiveClientData(tt.args.conn)
		})
	}
}

func Test_sendServerData(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendServerData(tt.args.conn)
		})
	}
}
