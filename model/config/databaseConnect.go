package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DatabaseConnect() (*sql.DB, error) {

	db, erro := sql.Open(PostgresDriver, DataSourceName)

	if erro != nil {
		return nil, erro
	}

	erro = db.Ping()
	if erro != nil {
		db.Close()
		return nil, erro
	}

	fmt.Println("Databse Connected")

	return db, nil
}
