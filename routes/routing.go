package routes

import (
	"fmt"
	"messagingApp/controllers"
	"messagingApp/websocket"
	"net/http"
)

func Routing(port string, hub *websocket.Hub) {
	http.HandleFunc("/messages/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetChatRoom(w, r)
	})

	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectSocket(w, r, hub)
	})

	fmt.Println("conn")

	http.ListenAndServe(port, nil)
}
