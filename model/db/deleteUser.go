package db

import (
	"database/sql"
	"errors"
)

func DeleteUserInDatabase(db *sql.DB, username string) error {

	exists, err := UserExists(db, username)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("user not exists")
	}

	_, err = db.Exec("DELETE from users WHERE username = $1", username)
	if err != nil {
		return err
	}

	return err
}
