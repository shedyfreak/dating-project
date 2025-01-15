package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetEvents(db *sqlx.DB) (events []Event) {
	events = make([]Event, 0)
	err := db.Select(&events, "SELECT * FROM events ORDER BY date ASC LIMIT 3")
	if err != nil {
		log.Fatal("cannot load events reason --> ", err.Error())
	}
	return events
}
