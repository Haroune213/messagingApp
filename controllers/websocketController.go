package controllers

import (
	"messagingApp/websocket"
	"net/http"
)

func ConnectSocket(w http.ResponseWriter, r *http.Request, hub *websocket.Hub) {

	websocket.WebSocket(w, r, hub)
}
