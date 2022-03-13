package orgsocket

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	connected     bool
	connectedSync sync.Mutex
)

func TriggerClient() {
	fmt.Println("Client started...")
	for {
		connectedSync.Lock()
		alreadyConnected := connected
		connectedSync.Unlock()
		if !alreadyConnected {
			conn, err := net.Dial("tcp", "127.0.0.1:8000")
			if err != nil {
				fmt.Println(err.Error())
				time.Sleep(time.Duration(5) * time.Second)
				continue
			}
			fmt.Println(conn.RemoteAddr().String() + ": connected")
			connectedSync.Lock()
			connected = true
			connectedSync.Unlock()
			go sendClientData(conn)
			go receiveServerData(conn)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func receiveServerData(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": disconnected")
			err1 := conn.Close()
			if err1 != nil {
				return
			}
			connectedSync.Lock()
			connected = false
			connectedSync.Unlock()
			fmt.Println(conn.RemoteAddr().String() + ": end receiving data")
			return
		}
		fmt.Print(conn.RemoteAddr().String() + ": received " + message)
	}
}

func sendClientData(conn net.Conn) {
	i := 0
	for {
		_, err := fmt.Fprintf(conn, strconv.Itoa(i)+". data from client\n")
		i++
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": end sending data")
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
