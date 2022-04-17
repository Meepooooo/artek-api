package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS test (
    id SERIAL PRIMARY KEY
);
`

func Database(user string, password string, dbName string) (*sqlx.DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbName)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
