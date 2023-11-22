package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	message := []byte("Hello, WebSocket Server!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal(err)
	}

	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received: %s\n", response)
}
