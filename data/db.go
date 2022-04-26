package data

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS rooms(
	id INTEGER PRIMARY KEY,
	date DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS teams(
    id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	room_id INTEGER NOT NULL,
	FOREIGN KEY(room_id) REFERENCES rooms(id)
);

CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY,
	team_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	role INTEGER NOT NULL
);
`

func Database(config config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", config.DSN)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
