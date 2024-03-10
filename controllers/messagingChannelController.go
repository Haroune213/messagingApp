package controllers

import (
	"fmt"
	"html/template"
	"messagingApp/database"
	"messagingApp/middlewares"
	"messagingApp/models"
	"net/http"
)

func GetChatRoom(w http.ResponseWriter, r *http.Request, url string) {
	_, id := middlewares.FilterUser(w, r)
	fmt.Println("user id: ", id)
	switch id {
	case 0:
		http.Redirect(w, r, "http://localhost:8000/login", http.StatusSeeOther)
	default:
		fmt.Println("url: ", url, " , id: ", id)
		val, exist := models.GetChannel(url, id)
		fmt.Println(val, exist)
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

func GetOrCreateChatRoom(w http.ResponseWriter, r *http.Request, url string) {
	_, id := middlewares.FilterUser(w, r)

	if id != 0 {
		target, err := database.GetUserByName(url)
		fmt.Println("target id: ", target.ID)

		if err != nil {
			fmt.Println(err)
		}

		link := models.CreateChannel(id, target.ID)
		fmt.Println("link: ", link)

		newUrl := "http://localhost:8000/web/" + link

		http.Redirect(w, r, newUrl, http.StatusSeeOther)

	}

}
