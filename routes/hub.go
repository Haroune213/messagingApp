package routes

import (
	"fmt"
	"sync"
)

type Message struct {
	Client_id string `json:"clientID"`
	Message   string `json:"text"`
}

type Hub struct {
	sync.RWMutex

	clients    map[*Client]bool //check if clients in a chat are connected or not
	messages   []*Message
	brodcast   chan *Message
	register   chan *Client
	unregister chan *Client
}

func CreateHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		brodcast:   make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()

		case client := <-h.unregister:
			h.Lock()
			if _, exist := h.clients[client]; exist {
				close(client.send)
				delete(h.clients, client)
			}
			h.Unlock()

		case msg := <-h.brodcast:
			h.RLock()

			h.messages = append(h.messages, msg)
			fmt.Println("msg: " + string(msg.Message))

			for client := range h.clients {
				select {

				case client.send <- getMessageTemplate(msg):

				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.RUnlock()

		}
	}
}
