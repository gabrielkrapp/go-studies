package db

import (
	"database/sql"
)

func SaveUser(db *sql.DB, username, hashedPassword string) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	return err
}
