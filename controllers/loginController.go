package controllers

import (
	"messagingApp/models"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, _ := models.GetUser(r.FormValue("email"), r.FormValue("password"))
	if user.ID == 0 {
		http.ServeFile(w, r, "./templates/errorLogin.html")
	} else {
		w.Header().Set("HX-Redirect", "http://localhost:8000/")
		w.WriteHeader(http.StatusOK)
	}

}
