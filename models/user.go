package models

import "github.com/dakshing/event-booking-rest-api-go/db"

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

	err = stmt.QueryRow(u.Username, u.Password).Scan(&u.ID)
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
