package database

type Event struct {
	ID          uint    `db:"id"`
	Name        string  `db:"name"`
	Location    string  `db:"location"`
	Date        string  `db:"date"`
	Time        string  `db:"time"`
	Price       float32 `db:"price"`
	Description string  `db:"description"`
	Icon        string  `db:"icon"`
}
