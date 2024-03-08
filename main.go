package main

import (
	"messagingApp/routes"
	"messagingApp/websocket"
)

func main() {
	hub := websocket.CreateHub()
	go hub.Run()

	routes.Routing(":8000", hub)

}
