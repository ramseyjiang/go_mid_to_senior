package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is return msg
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
}

// Start is before the project runs, the program starts > go Manager.Start ()
func (manager *ClientManager) Start() {
	for {
		log.Println("< --- pipeline communication -- >")
		select {
		case conn := <-Manager.Register:
			log.Printf("new user joined in %v", conn.ID)
			Manager.Clients[conn.ID] = conn
			jsonMessage, _ := json.Marshal(&Message{Content: "Successful connection to socket service"})
			conn.Send <- jsonMessage
		case conn := <-Manager.Unregister:
			log.Printf("user left %v", conn.ID)

			if _, ok := Manager.Clients[conn.ID]; ok {
				jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected"})
				conn.Send <- jsonMessage
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case message := <-Manager.Broadcast:
			MessageStruct := Message{}
			_ = json.Unmarshal(message, &MessageStruct)

			for ID, conn := range Manager.Clients {
				if ID != creatID(MessageStruct.Recipient, MessageStruct.Sender) {
					continue
				}
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
		}
	}
}
func creatID(sID, rID string) string {
	return sID + "_" + rID
}
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()

		if err != nil {
			Manager.Unregister <- c
			_ = c.Socket.Close()

			break
		}

		log.Printf("message read to client: %s", string(message))

		Manager.Broadcast <- message
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message := <-c.Send:
			log.Printf("message sent to client: %s", string(message))
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		default:
			_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
	}
}

// ChatHandler socket connection middleware function: upgrade protocol, user authentication, user-defined information, etc
func ChatHandler(c *gin.Context) {
	sID := c.Query("sId") // sender ID
	rID := c.Query("rId") // receiver ID
	log.Println(sID, rID)

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// User information authentication can be added
	client := &Client{
		ID:     creatID(sID, rID),
		Socket: conn,
		Send:   make(chan []byte),
	}

	Manager.Register <- client

	go client.Read()
	go client.Write()
}
