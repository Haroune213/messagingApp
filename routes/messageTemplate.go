package routes

import (
	"bytes"
	"log"
	"text/template"
)

func getMessageTemplate(msg *Message) []byte {
	tmpl, err := template.ParseFiles("templates/receivedMessage.html")

	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	var renderedMessage bytes.Buffer

	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
}
