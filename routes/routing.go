package routes

import (
	"fmt"
	"messagingApp/controllers"
	"messagingApp/websocket"
	"net/http"
)

func Routing(port string, hub *websocket.Hub) {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetLogin(w, r)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetRegister(w, r)
	})

	http.HandleFunc("/messages/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetChatRoom(w, r)
	})

	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectSocket(w, r, hub)
	})

	fmt.Println("conn")

	http.ListenAndServe(port, nil)
}
