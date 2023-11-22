package main

import (
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/net/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "ws://localhost:8080/ws"
	ws, err := websocket.Dial(url, "", "http://localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var message string
			err := websocket.Message.Receive(ws, &message)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Received message: %s\n", message)
		}
	}()


	fmt.Println("Enter message:")

	go func() {
		for {
			var input string
			fmt.Scanln(&input)

			err := websocket.Message.Send(ws, input)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			fmt.Println("Interrupt received, closing connection...")
			err := websocket.Message.Send(ws, "Client is disconnecting...")
			if err != nil {
				fmt.Println("Error closing connection:", err)
			}
			return
		}
	}
}
