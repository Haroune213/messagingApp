package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("connection open")

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println(err)
			break
		}
	}
}
