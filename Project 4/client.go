package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Введите сообщение для отправки на сервер:")

	go receiveMessages(conn)

	// Отправка сообщений с консоли
	scanner := bufio.NewScanner(os.Stdin)

	// Обработка сигнала завершения
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:
			fmt.Println("Завершение работы клиента.")
			return
		default:
			if scanner.Scan() {
				message := scanner.Text()
				err := conn.WriteMessage(websocket.TextMessage, []byte(message))
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func receiveMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Received: %s\n", message)
	}
}
