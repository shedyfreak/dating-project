package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBinit() (*sqlx.DB, error) {
	// Connection string
	const connStr = "user=myuser password=mypassword dbname=mydatabase host=localhost port=5432 sslmode=disable"

	// Open connection
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Println("Error opening database connection:", err)
		return nil, err
	}

	// Verify the connection is alive
	err = db.Ping()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}
	return db, nil
}
