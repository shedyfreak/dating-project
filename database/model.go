package database

import (
	"time"
)

type Event struct {
	ID           uint     `db:"id"`
	Name         string   `db:"name"`
	Location     string   `db:"location"`
	Date         DateOnly `db:"date"` // Implemets String() for tmpl
	Time         TimeOnly `db:"time"` // Implemets String() for tmpl
	Price        float32  `db:"price"`
	Description  string   `db:"description"`
	Icon         string   `db:"icon"`
	MapsLink     string   `db:"maps_link"`
	MaxSubsCount int      `db:"max_subs_count"`
}

type DateOnly time.Time // Alias for time.Time to represent a date only
type TimeOnly time.Time // Alias for time.Time to represent time of day only

func (d DateOnly) String() string {
	return time.Time(d).Format("2006-01-02")
}

// Format TimeOfDay as a string
func (t TimeOnly) String() string {
	return time.Time(t).Format("15:04:05")
}
