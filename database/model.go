package database

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Event struct {
	ID              uint     `db:"id"`
	Name            string   `db:"name"`
	Location        string   `db:"location"`
	Date            DateOnly `db:"date"` // Implemets String() for tmpl
	Time            TimeOnly `db:"time"` // Implemets String() for tmpl
	Price           float32  `db:"price"`
	Description     string   `db:"description"`
	Icon            string   `db:"icon"`
	MapsLink        string   `db:"maps_link"`
	MaxSubsCount    int      `db:"max_subs_count"`
	LongDescription string   `db:"long_description"`
	MinAge          int      `db:"min_age"`
	MaxAge          int      `db:"max_age"`
}
type Subsriber struct {
	ID          uint     `db:"id"`
	FirstName   string   `db:"first_name"`
	FamilyName  string   `db:"family_name"`
	Birthday    DateOnly `db:"birthday"` // Implemets String() for tmpl
	Email       string   `db:"email"`    // Implemets String() for tmpl
	PhoneNumber string   `db:"phone_number"`
	Sex         string   `db:"sex"`
}

type EventSubscription struct {
	EventID     uint   `db:"event_id"`
	SubsriberID uint   `db:"subsriber_id"`
	SessionID   string `db:"session_id"`
	Paid        bool   `db:"paid"`
}

type DateOnly time.Time // Alias for time.Time to represent a date only
type TimeOnly time.Time // Alias for time.Time to represent time of day only

func (d DateOnly) String() string {
	return time.Time(d).Format(time.DateOnly)
}

// Format TimeOfDay as a string
func (t TimeOnly) String() string {
	return time.Time(t).Format(time.TimeOnly)
}

func (d DateOnly) Value() (driver.Value, error) {
	return driver.Value(d.String()), nil
}
func (d *DateOnly) Scan(value interface{}) error {
	val, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("invalid date value")
	}
	parsedTime, err := time.Parse(time.DateOnly, val.String()[:10])
	if err != nil {
		return fmt.Errorf("couldn't parse date only value")
	}
	*d = DateOnly(parsedTime)
	return nil
}
func (d TimeOnly) Value() (driver.Value, error) {
	return driver.Value(d.String()), nil
}
func (d *TimeOnly) Scan(value interface{}) error {
	val, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("invalid date value")
	}
	parsedTime, err := time.Parse(time.TimeOnly, val.String()[11:19])
	if err != nil {
		return fmt.Errorf("couldn't parse date only value")
	}
	*d = TimeOnly(parsedTime)
	return nil
}
