package models

import (
	"time"

	"github.com/dakshing/event-booking-rest-api-go/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
	if err != nil {
		return err
	}

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=$1"
	var e Event
	err := db.DB.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events 
	SET name=$1, description=$2, location=$3, datetime=$4, user_id=$5 
	WHERE id=$6
	`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	return err
}

func DeleteEvent(id int64) error {
	query := "DELETE FROM events WHERE id=$1"
	_, err := db.DB.Exec(query, id)
	return err
}
