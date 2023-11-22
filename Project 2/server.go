package main

import (
	"fmt"
	"net/http"
	
	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	defer ws.Close()
	fmt.Println("Client Connected")

	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Received message: %s\n", message)

		err = websocket.Message.Send(ws, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(echoHandler))
	fmt.Println("WebSocket Echo Server is listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}