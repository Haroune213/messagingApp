package main

import (
	"log"
	"messagingApp/routes"
	"net/http"
)

func main() {

	http.HandleFunc("/web", routes.WebSocket)
	log.Fatal(http.ListenAndServe(":8008", nil))
}
