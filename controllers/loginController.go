package controllers

import "net/http"

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}
