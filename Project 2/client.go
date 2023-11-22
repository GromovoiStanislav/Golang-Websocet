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
		fmt.Println("Error connecting to the server:", err)
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
				fmt.Println("Error receiving message:", err)
				return
			}
			fmt.Printf("Received message: %s\n", message)
		}
	}()

	fmt.Println("Client connected to the server. Type 'exit' to disconnect.")
	fmt.Println("Enter message: ")

	go func() {
		for {
			var input string
			
			fmt.Scanln(&input)

			if input == "exit" {
				fmt.Println("Disconnecting...")
				close(done)
				return
			}

			err := websocket.Message.Send(ws, input)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		}
	}()

	<-done
}