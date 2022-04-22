package data

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS test (
    id INTEGER PRIMARY KEY
);
`

func Database(config config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", config.DSN)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
