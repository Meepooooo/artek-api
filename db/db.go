package db

import (
	"fmt"

	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS test (
    id SERIAL PRIMARY KEY
);
`

func Database(config config.DBConfig) (*sqlx.DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s", config.User, config.Password, config.Name)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
