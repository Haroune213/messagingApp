package api

import (
	"bytes"
	"html/template"
	"log"
	"messagingApp/database"
	"messagingApp/middlewares"
	"messagingApp/structs"
	"net/http"
)

func SideBarContact(w http.ResponseWriter, r *http.Request) {
	_, _, id := middlewares.VerifyJWT(w, r)
	channels := database.GetChannelsList(id)

	for _, channel := range channels {
		channel.Target_user, _ = database.GetUserById(id)
		displayContactTemplate(channel, w)
	}
}

func displayContactTemplate(channel structs.Message_channel, w http.ResponseWriter) []byte {
	tmpl, err := template.ParseFiles("templates/contactChannel.html")
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	var renderedChannel bytes.Buffer

	err = tmpl.Execute(w, channel)

	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedChannel.Bytes()
}
