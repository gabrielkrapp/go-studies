package db

import (
	"database/sql"
	"errors"
)

func EditUserPasswordInDatabase(db *sql.DB, username, newPassword string) error {

	exists, err := UserExists(db, username)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("user not exists")
	}

	_, err = db.Exec("UPDATE users SET password = $1 WHERE username = $2", newPassword, username)
	if err != nil {
		return err
	}

	return err
}
