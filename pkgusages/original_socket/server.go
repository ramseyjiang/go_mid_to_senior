package orgsocket

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

func TriggerServer() {
	fmt.Println("Server started...")
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting original_socket server: " + err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error listening to client: " + err.Error())
			continue
		}
		fmt.Println(conn.RemoteAddr().String() + ": client connected")
		go receiveClientData(conn)
		go sendServerData(conn)
	}
}

func sendServerData(conn net.Conn) {
	i := 0
	for {
		_, err := fmt.Fprintf(conn, strconv.Itoa(i)+". data from server\n")
		i++
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": end sending data")
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func receiveClientData(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": client disconnected")
			err1 := conn.Close()
			if err1 != nil {
				return
			}
			fmt.Println(conn.RemoteAddr().String() + ": end receiving data")
			return
		}
		fmt.Print(conn.RemoteAddr().String() + ": received " + message)
	}
}
