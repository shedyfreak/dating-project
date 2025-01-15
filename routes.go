package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter(h handler) {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	// Define routes
	r.HandleFunc("/", h.homeHandler).Methods("GET")
	r.HandleFunc("/api/events", h.eventsHandler).Methods("GET")
	r.HandleFunc("/api/event-opts", h.eventsOptions).Methods("GET")
	r.HandleFunc("/api/register", h.registrationHandler).Methods("POST")
	//http.Handle("/", http.FileServer(http.Dir("stripe-public")))
	r.HandleFunc("/api/success", h.success).Methods("GET")
	r.HandleFunc("/api/cancel", h.success).Methods("GET")

	// Start the server
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("cannot init server, reason -> %s", err.Error())
		panic("db failure")
	}
}
