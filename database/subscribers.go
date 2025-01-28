package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// SaveSubscribe saves subscriber to db and return his ID
func SaveSubscriber(db *sqlx.DB, sub Subsriber) (err error, id uint) {
	query := ` INSERT INTO subscribers (first_name, family_name, birthday, email, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err = db.Get(&sub.ID, query, sub.FirstName, sub.FamilyName, sub.Birthday, sub.Email, sub.PhoneNumber)
	if err != nil {
		msg := "Failed to insert data into subs"
		log.Println(msg, err.Error())
		return
	}
	return nil, sub.ID
}
func SaveTempSubscription(db *sqlx.DB, eventID uint, subscriberID uint, sessionID string) (err error) {
	query := ` INSERT INTO event_subscriptions (event_id, subscriber_id, session_id)
		VALUES ($1, $2, $3)`
	_, err = db.Exec(query, eventID, subscriberID, sessionID)
	if err != nil {
		log.Println("issue with subscribing", err)
		return err
	}

	return nil
}

func SuccessPaidSubscription(db *sqlx.DB, sessionID string) (sub Subsriber, event Event, err error) {
	query := `UPDATE event_subscriptions 
          SET paid = true 
          WHERE session_id = $1 
          RETURNING subscriber_id,, event_id`

	var subscriberID int
	var eventID int
	err = db.QueryRow(query, sessionID).Scan(&subscriberID, eventID)
	if err != nil {
		log.Println("issue with subscribing:", err)
		return
	}

	var allSubs []Subsriber
	err = db.Get(&allSubs, "SELECT * FROM subscribers WHERE id=$1", subscriberID)
	if err != nil {
		log.Println("issue with geting successfful subscriber:", err)
		return
	}

	var allEvs []Event
	err = db.Select(&allEvs, "SELECT * FROM subscribers WHERE id=$1", eventID)
	if err != nil {
		log.Println("issue with geting successfful subscriber:", err)
		return
	}
	event = allEvs[0]
	sub = allSubs[0]
	return
}
