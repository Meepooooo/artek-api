package db

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE sessions(
    id INTEGER PRIMARY KEY
);

CREATE TABLE session_users(
    id INTEGER PRIMARY KEY,
    session_id INTEGER NOT NULL,
    user_name TEXT NOT NULL,
    user_role INTEGER NOT NULL,
	FOREIGN KEY(session_id) REFERENCES sessions(id)
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
