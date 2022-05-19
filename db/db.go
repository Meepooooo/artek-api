package db

import (
	"database/sql"

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
	name TEXT NOT NULL,
	role INTEGER NOT NULL,
	team_id INTEGER NOT NULL,
	FOREIGN KEY(team_id) REFERENCES teams(id)
);
`

func Database(location string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
