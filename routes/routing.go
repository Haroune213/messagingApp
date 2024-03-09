package routes

import (
	"fmt"
	"messagingApp/controllers"
	"messagingApp/websocket"
	"net/http"
	"strings"
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

	http.HandleFunc("/web/", func(w http.ResponseWriter, r *http.Request) {
		id := extractID(r.URL.Path)
		fmt.Println(id)
		controllers.GetChatRoom(w, r)
	})

	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectSocket(w, r, hub)
	})

	http.ListenAndServe(port, nil)
}

func extractID(path string) string {
	parts := strings.Split(path, "/")

	for i, part := range parts {
		if part == "web" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
