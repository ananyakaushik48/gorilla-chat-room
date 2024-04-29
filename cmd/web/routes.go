package main

import (
	"net/http"
	"ws/internal/handlers"

	"github.com/bmizerany/pat"
)


//  Using pat to return the Home page handler
func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	// Serving static JS for the webserver
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
