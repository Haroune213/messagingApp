package controllers

import (
	"fmt"
	"messagingApp/middlewares"
	"messagingApp/models"
	"net/http"
)

func GetRegister(w http.ResponseWriter, r *http.Request) {
	_, id := middlewares.FilterUser(w, r)

	switch id {
	case 0:
		http.ServeFile(w, r, "./templates/register.html")
	default:
		http.Redirect(w, r, "http://localhost:8000/", http.StatusSeeOther)
	}

}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	_, id := middlewares.FilterUser(w, r)
	if id != 0 {

		id, exist, err := models.CreateUser(r.FormValue("email"), r.FormValue("username"), r.FormValue("password"))

		if err != nil || id == 0 || !exist {
			fmt.Println("err: ", err, " exist: ", exist, " id: ", id)
			http.ServeFile(w, r, "./templates/errorRegister.html")
		} else {
			w.Header().Set("HX-Redirect", "http://localhost:8000/")
			w.WriteHeader(http.StatusOK)
		}
	}
}
