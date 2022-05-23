package db

type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoomID int    `json:"roomId,omitempty"`
	Users  []User `json:"users,omitempty"`
}

func (d *DB) GetTeam(id int) (team Team, err error) {
	err = d.db.QueryRow("SELECT id, name, room_id FROM teams WHERE id = ?;", id).Scan(&team.ID, &team.Name, &team.RoomID)
	if err != nil {
		return Team{}, err
	}

	rows, err := d.db.Query("SELECT id, name, role FROM users WHERE team_id = ?;", id)
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

func (d *DB) CreateTeam(name string, roomID int) (int, error) {
	var exists int
	err := d.db.QueryRow("SELECT 1 FROM rooms WHERE id = ?;", roomID).Scan(&exists)
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
