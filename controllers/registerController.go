package controllers

import "net/http"

func GetRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/register.html")
}
