package controllers

import (
	"messagingApp/middlewares"
	"net/http"
)

func GetChatRoom(w http.ResponseWriter, r *http.Request) {
	_, id := middlewares.FilterUser(w, r)

	switch id {
	case 0:
		http.Redirect(w, r, "http://localhost:8000/login", http.StatusSeeOther)
	default:
		//val := models.CreateChannel(1, 2)
		http.ServeFile(w, r, "./templates/index.html")

	}
}
