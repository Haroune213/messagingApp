package routes

import (
	"fmt"
	"messagingApp/api"
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

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			api.SearchUser(r.FormValue("username"), w)
		}
	})

	http.HandleFunc("/create/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("active")
		if r.Method == "GET" {

			url := extractID("create", r.URL.Path)
			fmt.Println(url)
			controllers.GetOrCreateChatRoom(w, r, url)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetHomePage(w, r)

	})

	http.HandleFunc("/web/", func(w http.ResponseWriter, r *http.Request) {
		url := extractID("web", r.URL.Path)
		controllers.GetChatRoom(w, r, url)

	})

	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectSocket(w, r, hub)
	})

	http.ListenAndServe(port, nil)
}

func extractID(prefix string, path string) string {
	parts := strings.Split(path, "/")

	for i, part := range parts {
		if part == prefix && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
