package api

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"messagingApp/database"
	"messagingApp/structs"
	"net/http"
)

//this api is used to search for users

func SearchUser(user string, w http.ResponseWriter) {
	users := database.GetUsersByName(user)
	for _, usr := range users {
		displayUserTemplate(usr, w)
	}
}

func displayUserTemplate(usr structs.User, w http.ResponseWriter) []byte {
	tmpl, err := template.ParseFiles("templates/searchResult.html")
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	fmt.Println("usr:", usr.Username)

	var renderedUsr bytes.Buffer

	err = tmpl.Execute(w, usr)

	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedUsr.Bytes()
}
