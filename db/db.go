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

func Database(location string) (*DB, error) {
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

type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoomID int    `json:"roomId,omitempty"`
	Users  []User `json:"users"`
}

func (d *DB) GetTeam(id int) (team Team, err error) {
	err = d.db.QueryRow("SELECT id, name, room_id FROM teams WHERE id = ?;", id).Scan(&team.ID, &team.Name, &team.RoomID)
	if err != nil {
		return Team{}, err
	}

	rows, err := d.db.Query("SELECT id, name, role FROM users WHERE team_id = ?", id)
	if err != nil {
		return Team{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
			return Team{}, err
		}
		team.Users = append(team.Users, user)
	}

	return team, nil
}

func (d *DB) CreateTeam(name string, roomID int) (id int, err error) {
	var exists int
	err = d.db.QueryRow("SELECT 1 FROM rooms WHERE id = ?;", roomID).Scan(&exists)
	if err != nil {
		return 0, err
	}

	res, err := d.db.Exec("INSERT INTO teams(name, room_id) VALUES(?, ?);", name, roomID)
	if err != nil {
		return 0, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(rowID), nil
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Role   int    `json:"role"`
	TeamID int    `json:"teamId,omitempty"`
}
