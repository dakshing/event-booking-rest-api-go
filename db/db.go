package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Global database connection pool

func InitDB() {
	connStr := "user=user password=password dbname=event_booking sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMP NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		log.Println("Error creating events table:", err)
		panic("Could not create events table.")
	}
}
