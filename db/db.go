package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS rooms(
	id   INTEGER  PRIMARY KEY,
	time DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS teams(
	id      INTEGER PRIMARY KEY,
	name    TEXT    NOT NULL,
	room_id INTEGER NOT NULL,
	water   INTEGER NOT NULL,
	food    INTEGER NOT NULL,
	oxygen  INTEGER NOT NULL,
	spirit  INTEGER NOT NULL,
	fuel    INTEGER NOT NULL,
	FOREIGN KEY(room_id) REFERENCES rooms(id)
);

CREATE TABLE IF NOT EXISTS users(
	id      INTEGER PRIMARY KEY,
	name    TEXT    NOT NULL,
	role    INTEGER NOT NULL,
	team_id INTEGER NOT NULL,
	FOREIGN KEY(team_id) REFERENCES teams(id)
);
`

// Open opens a database specified by its location.
func Open(location string) (*DB, error) {
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (d *DB) Ping() error {
	return d.db.Ping()
}
