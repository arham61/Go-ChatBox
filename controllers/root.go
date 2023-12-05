package controllers

import (
	"github.com/gorilla/websocket"
	"chatbox/structures"
	"net/http"
	"fmt"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan structures.MessageWithSender)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Chat Room!")
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg structures.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}

	
		broadcast <- structures.MessageWithSender{Sender: conn, Message: msg}
	}
}

func HandleMessages() {
	for {
		msgWithSender := <-broadcast
		sender := msgWithSender.Sender
		msg := msgWithSender.Message

		for client := range clients {
			if client == sender {
				continue
			}

			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}