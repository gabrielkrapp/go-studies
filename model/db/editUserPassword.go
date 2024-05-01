package db

import (
	"database/sql"
)

func EditUserPasswordInDatabase(db *sql.DB, username, newPassword string) error {
	_, err := db.Exec("UPDATE users SET password = $1 WHERE username = $2", newPassword, username)
	return err
}
