package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/perso-proj/database"
	"github.com/perso-proj/templates"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var GlobalDB *sqlx.DB

type handler struct {
	db *sqlx.DB
}

func (h *handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, templates.HomeTempl)
}

func (h *handler) eventsHandler(w http.ResponseWriter, r *http.Request) {
	events := database.GetEvents(h.db)
	res := templates.ParseEvents(events)

	fmt.Fprintf(w, res)
}

func main() {

	GlobalDB, err := database.DBinit()
	if err != nil {
		log.Fatal("cannot init database, reason -> %s", err.Error())
		panic("db failure")
	}

	hdler := handler{db: GlobalDB}
	// // Query the database
	// rows, err := GlobalDBb.Query("SELECT id, name FROM events")
	// if err != nil {
	// 	log.Fatal("Error executing query:", err)
	// }
	// defer rows.Close()

	// // Iterate over rows and display the result
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	err = rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal("Error scanning row:", err)
	// 	}
	// 	fmt.Printf("ID: %d, Name: %s\n", id, name)
	// }

	// // Check for any error encountered during iteration
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal("Error iterating over rows:", err)
	// }
	initRouter(hdler)

}

func initRouter(h handler) {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", h.homeHandler).Methods("GET")
	r.HandleFunc("/api/events", h.eventsHandler).Methods("GET")

	// Start the server
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("cannot init database, reason -> %s", err.Error())
		panic("db failure")
	}
}
