package gorillaws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/**
Package net/http is used for serving raw HTTP connections.
Regular HTTP server and add a socketHandler() function to handle the websocket logic are in server.go.

Gorilla converts these raw HTTP connections into a stateful websocket connection, using a connection upgrade.
This server simply echoes any incoming websocket messages back to the client.
*/

// use default options, it is a variable to help us convert any incoming HTTP connection into websocket protocol, via upgrade.Upgrade()
var upgrade = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgrade:", err)
		return
	}

	defer func(conn *websocket.Conn) {
		err2 := conn.Close()
		if err2 != nil {
			return
		}
	}(conn)

	// The event loop, The server reads messages using conn.ReadMessage() and writes them back using conn.WriteMessage()
	for {
		messageType, message, err1 := conn.ReadMessage() // using conn.ReadMessage() to read messages.
		if err1 != nil {
			log.Println("Error during message reading:", err1)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message) // using conn.WriteMessage() to write messages back.
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Index Page")
	if err != nil {
		return
	}
	fmt.Println(r)
}

// TriggerServer Send requests to localhost:8080/ and localhost:8080/socket, it will show differences print in terminal.
func TriggerServer() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
