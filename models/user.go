package models

import (
	"github.com/dakshing/event-booking-rest-api-go/db"
	"github.com/dakshing/event-booking-rest-api-go/utils"
)

type User struct {
	ID       int64
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (username, password) 
	VALUES ($1, $2) RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(u.Username, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return err
}

func GetUserByUsername(username string) (*User, error) {
	query := "SELECT * FROM users WHERE username=$1"
	var u User
	err := db.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
