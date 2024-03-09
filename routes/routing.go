package routes

import (
	"messagingApp/controllers"
	"messagingApp/websocket"
	"net/http"
)

func Routing(port string, hub *websocket.Hub) {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.GetLogin(w, r)
		}
		if r.Method == "POST" {
			controllers.PostLogin(w, r)
		}
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.GetRegister(w, r)
		}
		if r.Method == "POST" {
			controllers.PostRegister(w, r)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetChatRoom(w, r)
	})

	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectSocket(w, r, hub)
	})

	http.ListenAndServe(port, nil)
}
