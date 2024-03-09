package controllers

import (
	"messagingApp/middlewares"
	"messagingApp/models"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	_, id := middlewares.FilterUser(w, r)

	switch id {
	case 0:
		http.ServeFile(w, r, "./templates/login.html")

	default:
		http.Redirect(w, r, "http://localhost:8000/", http.StatusSeeOther)
	}

}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	_, id := middlewares.FilterUser(w, r)

	if id != 0 {

		r.ParseForm()

		user, _ := models.GetUser(r.FormValue("email"), r.FormValue("password"))
		if user.ID == 0 {
			http.ServeFile(w, r, "./templates/errorLogin.html")
		} else {

			middlewares.CreateJWT(user.ID, user.Username, user.Email, w, r) //this will create a JWTtoken stored in the cookies

			w.Header().Set("HX-Redirect", "http://localhost:8000/")
			w.WriteHeader(http.StatusOK)
		}
	}
}
