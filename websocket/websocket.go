package websocket

import (
	"net/http"
	"time"

	"github.com/google/uuid"
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
	user_id string
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
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	gen_id := uuid.New()

	client := &Client{
		user_id: gen_id.String(),
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte),
	}

	client.hub.register <- client

	go client.readMessage()
	go client.writeMessage()

}
