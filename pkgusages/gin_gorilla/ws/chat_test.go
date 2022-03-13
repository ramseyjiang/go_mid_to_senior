package ws

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestChatHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChatHandler(tt.args.c)
		})
	}
}

func TestClientManager_Start(t *testing.T) {
	type fields struct {
		Clients    map[string]*Client
		Broadcast  chan []byte
		Register   chan *Client
		Unregister chan *Client
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &ClientManager{
				Clients:    tt.fields.Clients,
				Broadcast:  tt.fields.Broadcast,
				Register:   tt.fields.Register,
				Unregister: tt.fields.Unregister,
			}
			manager.Start()
		})
	}
}

func TestClient_Read(t *testing.T) {
	type fields struct {
		ID     string
		Socket *websocket.Conn
		Send   chan []byte
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				ID:     tt.fields.ID,
				Socket: tt.fields.Socket,
				Send:   tt.fields.Send,
			}
			c.Read()
		})
	}
}

func TestClient_Write(t *testing.T) {
	type fields struct {
		ID     string
		Socket *websocket.Conn
		Send   chan []byte
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				ID:     tt.fields.ID,
				Socket: tt.fields.Socket,
				Send:   tt.fields.Send,
			}
			c.Write()
		})
	}
}

func Test_creatID(t *testing.T) {
	type args struct {
		sID string
		rID string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := creatID(tt.args.sID, tt.args.rID); got != tt.want {
				t.Errorf("creatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
