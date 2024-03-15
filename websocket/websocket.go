package websocket

import (
	"fmt"
	"messagingApp/middlewares"
	"net/http"
)

func WebSocket(w http.ResponseWriter, r *http.Request, hub *Hub, link string) {
	_, _, id := middlewares.VerifyJWT(w, r)
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	client := &Client{
		user_id: id,
		link:    link,
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte),
	}

	fmt.Println("Client: ", client.user_id, " link: ", client.link)

	client.hub.register <- client

	go client.readMessage(link)
	go client.writeMessage()

}
