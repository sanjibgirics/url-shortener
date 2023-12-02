package main

import (
	"github.com/gorilla/mux"
	"github.com/sanjibgirics/url-shortener/pkg"
	"log"
	"net/http"
)

func main() {
	log.Println("Creating url shortener endpoints.")
	// Create router
	r := mux.NewRouter()

	// Register all the endpoints to the router
	pkg.RegisterRoutes(r)

	// Start listening and serving
	log.Println("Listening and serving url shortener endpoints . . .")
	http.ListenAndServe(":8082", r)
}
