package db

import (
	"database/sql"
)

func GetListOfUsers(db *sql.DB) ([]string, error) {
	var users []string

	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		users = append(users, username)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
