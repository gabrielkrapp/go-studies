package db

import (
	"database/sql"
	"errors"
)

func RegisterUser(db *sql.DB, username, hashedPassword string) error {

	exists, err := UserExists(db, username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		return err
	}

	return err
}
