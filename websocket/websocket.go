package websocket

import (
	"messagingApp/middlewares"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pingPeriod     = (pongWait * 9) / 10
	pongWait       = 10 * time.Second
	maxMessageSize = 512
	writeWait      = 10 * time.Second
)

type WsMessage struct {
	Message string      `json:"message"`
	Headers interface{} `json:"HEADERS"`
}

type Client struct {
	user_id int
	hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(w http.ResponseWriter, r *http.Request, hub *Hub) {
	_, _, id := middlewares.VerifyJWT(w, r)
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	client := &Client{
		user_id: id,
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte),
	}

	client.hub.register <- client

	go client.readMessage()
	go client.writeMessage()

}
