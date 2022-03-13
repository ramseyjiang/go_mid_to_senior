package gorillaws

import (
	"testing"

	"github.com/gorilla/websocket"
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

func Test_receiveHandler(t *testing.T) {
	type args struct {
		connection *websocket.Conn
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiveHandler(tt.args.connection)
		})
	}
}
