package controllers

import "net/http"

func GetChatRoom(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")

}
