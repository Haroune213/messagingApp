package controllers

import (
	"fmt"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Printf("Email:\n%s", r.FormValue("email"))
	fmt.Printf("Password:\n%s", r.FormValue("password"))

}
