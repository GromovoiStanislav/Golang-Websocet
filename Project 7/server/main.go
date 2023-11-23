package main

import (
	"fmt"
	"time"


	"simple-webcoket/server/ws"
)

func main() {
	server := ws.StartServer(messageHandler)

	for {
		server.WriteMessage([]byte("Hello"))
		time.Sleep(time.Second)
	}
}

func messageHandler(message []byte) {
	fmt.Println(string(message))
}