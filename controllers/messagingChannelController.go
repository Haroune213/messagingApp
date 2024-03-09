package controllers

import (
	"fmt"
	"messagingApp/middlewares"
	"messagingApp/models"
	"net/http"
)

func GetChatRoom(w http.ResponseWriter, r *http.Request, url string) {
	_, id := middlewares.FilterUser(w, r)

	switch id {
	case 0:
		http.Redirect(w, r, "http://localhost:8000/login", http.StatusSeeOther)
	default:
		val := models.GetChannel(url, id)

		fmt.Println(val)
		http.ServeFile(w, r, "./templates/index.html")

	}
}
