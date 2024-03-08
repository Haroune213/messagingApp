package controllers

import (
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("email") != "123@mail.com" || r.FormValue("password") != "123" {
		http.ServeFile(w, r, "./templates/errorLogin.html")
	}
	if r.FormValue("email") == "123@mail.com" && r.FormValue("password") == "123" {
		w.Header().Set("HX-Redirect", "http://localhost:8000/")
		w.WriteHeader(http.StatusOK)

	}

}
