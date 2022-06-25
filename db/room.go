package db

import "time"

// Room represents a quest room/instance.
// Teams can connect to a room to play in the same environment
// and be able to communicate, trade, etc.
type Room struct {
	ID    int       `json:"id"`
	Time  time.Time `json:"time"`
	Teams []Team    `json:"teams"`
}

// GetRoom gets a room specified by its ID.
// If no room is found, GetRoom will return an empty Room and sql.ErrNoRows.
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

	room.Teams = make([]Team, 0)

	for rows.Next() {
		var team Team
		if err = rows.Scan(&team.ID, &team.Name); err != nil {
			return Room{}, err
		}
		room.Teams = append(room.Teams, team)
	}

	return room, nil
}

// CreateRoom creates a new room with the current UTC time.
func (d *DB) CreateRoom() (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO ROOMS(time) VALUES(DATETIME('now'));")
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
