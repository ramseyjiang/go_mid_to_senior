package ws

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const times = 10

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TextAPI webSocket returns text format
func TextAPI(c *gin.Context) {
	// Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
	}

	defer func(ws *websocket.Conn) {
		err1 := ws.Close()
		if err1 != nil {
			return
		}
	}(ws)

	// Read data in ws
	mt, message, err2 := ws.ReadMessage()
	if err2 != nil {
		log.Println("error read message")
	}

	// Write ws data, pong 10 times
	var count = 0

	for {
		count++
		if count > times {
			break
		}

		message = []byte(string(message) + " " + strconv.Itoa(count))
		err = ws.WriteMessage(mt, message)

		if err != nil {
			log.Println("error write message: " + err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}

// JSONAPI webSocket returns json format
func JSONAPI(c *gin.Context) {
	// Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
	}

	defer func(ws *websocket.Conn) {
		err1 := ws.Close()
		if err1 != nil {
			return
		}
	}(ws)

	var data struct {
		A string `json:"a"`
		B int    `json:"b"`
	}

	// Read data in ws
	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
	}

	// Write ws data, pong 10 times
	var count = 0

	for {
		count++
		if count > times {
			break
		}

		err = ws.WriteJSON(struct {
			A string `json:"a"`
			B int    `json:"b"`
			C int    `json:"c"`
		}{
			A: data.A,
			B: data.B,
			C: count,
		})
		if err != nil {
			log.Println("error write json: " + err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}
