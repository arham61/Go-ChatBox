package structures

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MessageWithSender struct {
	Sender  *websocket.Conn
	Message Message
}