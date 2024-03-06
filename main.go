package main

import (
	"fmt"
	"log"
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
	http.HandleFunc("/web", routes.WebSocket)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
