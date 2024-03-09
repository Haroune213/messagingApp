package controllers

import (
	"html/template"
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
		val, exist := models.GetChannel(url, id)
		if exist {
			tmpl, err := template.ParseFiles("./templates/index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			tmpl.Execute(w, val)
		} else {
			http.Redirect(w, r, "http://localhost:8000/404", http.StatusSeeOther)

		}
	}
}
