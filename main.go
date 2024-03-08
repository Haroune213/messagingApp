package main

import (
	"messagingApp/database"
	"messagingApp/routes"
	"messagingApp/websocket"
)

func main() {
	hub := websocket.CreateHub()
	go hub.Run()
	database.OpenDB()
	routes.Routing(":8000", hub)

}
