package db

import "time"

type Room struct {
	ID    int       `json:"id"`
	Time  time.Time `json:"time"`
	Teams []Team    `json:"teams"`
}

func (d *DB) GetRoom(id int) (room Room, err error) {
	err = d.db.QueryRow("SELECT id, time FROM rooms WHERE id = ?;", id).Scan(&room.ID, &room.Time)
	if err != nil {
		return Room{}, err
	}

	rows, err := d.db.Query("SELECT id, name FROM teams WHERE room_id = ?;", id)
	if err != nil {
		return Room{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var team Team
		if err = rows.Scan(&team.ID, &team.Name); err != nil {
			return Room{}, err
		}
		room.Teams = append(room.Teams, team)
	}

	return room, nil
}

func (d *DB) CreateRoom() (id int, err error) {
	res, err := d.db.Exec("INSERT INTO ROOMS(time) VALUES(DATETIME('now'));")
	if err != nil {
		return 0, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(rowID), nil
}
