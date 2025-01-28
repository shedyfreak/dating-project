package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetEvents(db *sqlx.DB) (events []Event, err error) {
	events = make([]Event, 0)
	err = db.Select(&events, "SELECT * FROM events ORDER BY date ASC LIMIT 3")
	if err != nil {
		log.Println("cannot load events reason --> ", err.Error())
		return
	}
	return
}

func GetEligibleBySex(db *sqlx.DB, id uint, sex string) (eligible bool, leftPlacesMale int, leftPlacesFemales int, err error) {
	type SexCount struct {
		Sex   string `json:"sex"`
		Count int    `json:"count"`
		Max   int    `json:"max_subs_count"`
	}

	query := `SELECT s.sex,
    	COUNT(*) AS count
		FROM 
    	event_subscriptions es
		JOIN 
    	subscribers s
		ON 
    	es.subscriber_id = s.id
		WHERE 
    	es.event_id = $1 -- Replace $1 with the event_id you're checking
		GROUP BY 
    	s.sex;`
	var count []SexCount
	err = db.Select(&count, query, id)
	if err != nil {
		log.Println("cannot load events reason --> ", err.Error())
		return
	}
	max_event := count[0].Max
	for _, sexCount := range count {
		if sexCount.Sex == "male" {
			leftPlacesMale = max_event/2 - sexCount.Count
		}
		if sexCount.Sex == "female" {
			leftPlacesFemales = max_event/2 - sexCount.Count
		}
	}
	if leftPlacesMale <= 0 && sex == "male" {
		return
	}
	if leftPlacesFemales <= 0 && sex == "female" {
		return
	}
	eligible = true
	return
}
