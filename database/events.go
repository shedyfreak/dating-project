package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetEvents(db *sqlx.DB) (events []Event) {
	events = make([]Event, 0)
	err := db.Select(&events, "SELECT * FROM events")
	if err != nil {
		log.Fatal("cannot load events reason --> %s", err.Error())
	}
	return events
}
