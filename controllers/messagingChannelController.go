package controllers

import (
	"fmt"
	"html/template"
	"messagingApp/api"
	"messagingApp/database"
	"messagingApp/middlewares"
	"messagingApp/models"
	"messagingApp/structs"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	_, id := middlewares.FilterUser(w, r)

	switch id {
	case 0:
		http.Redirect(w, r, "http://localhost:8000/login", http.StatusSeeOther)
	default:
		tmpl, err := template.ParseFiles("./templates/homepage.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		contacts := api.SideBarContact(id)

		pageData := &structs.ChannelPage{
			User:     structs.User{},
			Contacts: contacts,
		}

		tmpl.Execute(w, pageData)
	}
}

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

			contacts := api.SideBarContact(id)

			pageData := &structs.ChannelPage{
				User:     val,
				Link:     url,
				Contacts: contacts,
			}

			tmpl.Execute(w, pageData)
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
