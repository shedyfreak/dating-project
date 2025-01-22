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
	err = db.Get(&sub.ID, query, sub.FamilyName, sub.FamilyName, sub.Birthday, sub.Email, sub.PhoneNumber)
	if err != nil {
		msg := "Failed to insert data into subs"
		log.Fatalln(msg, err.Error())
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

func SuccessPaidSubscription(db *sqlx.DB, sessionID string) (err error) {
	query := ` UPDATE event_subscriptions SET paid = true where session_id = $1`
	_, err = db.Exec(query, sessionID)
	if err != nil {
		log.Println("issue with subscribing")
		return err
	}

	return nil
}
