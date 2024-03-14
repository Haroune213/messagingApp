package websocket

import (
	"messagingApp/middlewares"
	"net/http"
)

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
