package websocket

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pingPeriod     = (pongWait * 9) / 10
	pongWait       = 10 * time.Second
	maxMessageSize = 512
	writeWait      = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

type Message struct {
	Client_id int    `json:"clientID"`
	Message   string `json:"text"`
}

type Hub struct {
	sync.RWMutex
	url        string
	clients    map[*Client]bool //check if clients in a chat are connected or not
	messages   []*Message
	brodcast   chan *Message
	register   chan *Client
	unregister chan *Client
}
