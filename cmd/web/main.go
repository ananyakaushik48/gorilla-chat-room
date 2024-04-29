package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {
	// All the routes are here
	mux := routes()
	log.Println("Starting the channel listener")

	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")
	// This will server the routes at 8080
	_ = http.ListenAndServe(":8080", mux)
}
