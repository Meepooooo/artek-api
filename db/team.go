package db

var defaultResources = map[int]Resources{
	rolePlayer:  {2, 2, 2, 2},
	roleCaptain: {3, 3, 3, 3},
}

type Resources struct {
	Water  int `json:"Water"`
	Food   int `json:"Food"`
	Oxygen int `json:"Oxygen"`
	Spirit int `json:"Spirit"`
}

// Team represents a quest team.
// A team belongs to a room and has multiple users.
type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoomID int    `json:"roomId,omitempty"`
	Users  []User `json:"users"`
	Resources
}

// GetTeam gets a team specified by its ID.
// If no team is found, GetTeam will return an empty Team and sql.ErrNoRows.
func (d *DB) GetTeam(id int) (team Team, err error) {
	err = d.db.QueryRow(`SELECT id, name, room_id, water, food, oxygen, spirit
		FROM teams WHERE id = ?;`, id).Scan(&team.ID, &team.Name, &team.RoomID,
		&team.Water, &team.Food, &team.Oxygen, &team.Spirit)
	if err != nil {
		return Team{}, err
	}

	rows, err := d.db.Query("SELECT id, name, role FROM users WHERE team_id = ?;", id)
	if err != nil {
		return Team{}, err
	}
	defer rows.Close()

	team.Users = make([]User, 0)

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
			return Team{}, err
		}
		team.Users = append(team.Users, user)
	}

	return team, nil
}

// CreateTeam creates a new team with the specified name and room ID.
// If no room with the specified ID is found, CreateTeam will return 0 and sql.ErrNoRows.
func (d *DB) CreateTeam(name string, roomID int) (id int64, err error) {
	err = d.db.QueryRow("SELECT 1 FROM rooms WHERE id = ?;", roomID).Scan(new(byte))
	if err != nil {
		return 0, err
	}

	res, err := d.db.Exec(`INSERT INTO teams(name, room_id, water, food, oxygen, spirit)
		VALUES(?, ?, 0, 0, 0, 0);`, name, roomID)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *DB) SetTeamBalance(id int) (balance Resources, err error) {
	team, err := d.GetTeam(id)
	if err != nil {
		return Resources{}, err
	}

	for _, user := range team.Users {
		userResources := defaultResources[user.Role]

		balance.Water += userResources.Water
		balance.Food += userResources.Food
		balance.Oxygen += userResources.Oxygen
		balance.Spirit += userResources.Spirit
	}

	_, err = d.db.Exec(`UPDATE teams
		SET water = ?, food = ?, oxygen = ?, spirit = ?
		WHERE id = ?`,
		balance.Water, balance.Food, balance.Oxygen, balance.Spirit, id)
	if err != nil {
		return Resources{}, err
	}

	return balance, nil
}
