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

func GetEventByID(db *sqlx.DB, name string) (err error, id uint) {

	err = db.Select(&id, "SELECT id FROM events WHERE name = $1", name)
	if err != nil {
		log.Fatal("cannot load events reason --> ", err.Error())
		return
	}
	return nil, id
}
