package controllers

import (
	"fmt"
	"messagingApp/models"
	"net/http"
)

func GetRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/register.html")
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	id, exist, err := models.CreateUser(r.FormValue("email"), r.FormValue("username"), r.FormValue("password"))

	if err != nil || id == 0 || !exist {
		fmt.Println("err: ", err, " exist: ", exist, " id: ", id)
		http.ServeFile(w, r, "./templates/errorRegister.html")
	} else {
		w.Header().Set("HX-Redirect", "http://localhost:8000/")
		w.WriteHeader(http.StatusOK)
	}

}
