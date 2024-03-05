package main

import (
	"fmt"
	"messagingApp/routes"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the HTML page
		http.ServeFile(w, r, "./templates/index.html")
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/web", routes.WebSocket)
}
