package db

import (
	"database/sql"
)

type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoomID int    `json:"roomId"`
	Users  []User `json:"users"`
}

func GetTeam(db *sql.DB, id int) (team Team, err error) {
	err = db.QueryRow("SELECT id, name, room_id FROM teams WHERE id = ?;", id).Scan(&team.ID, &team.Name, &team.RoomID)
	if err != nil {
		return Team{}, err
	}

	rows, err := db.Query("SELECT id, name, role FROM users WHERE team_id = ?", id)
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

func CreateTeam(db *sql.DB, name string, roomID int) (id int, err error) {
	var exists int
	err = db.QueryRow("SELECT 1 FROM rooms WHERE id = ?;", roomID).Scan(&exists)
	if err != nil {
		return 0, err
	}

	res, err := db.Exec("INSERT INTO teams(name, room_id) VALUES(?, ?);", name, roomID)
	if err != nil {
		return 0, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(rowID), nil
}
