package db

// User represents a quest user.
// A user belongs to a team.
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Role   int    `json:"role"`
	TeamID int    `json:"teamId,omitempty"`
}

// CreateRoom creates a new room with the specified name, role and team ID.
// If no team with the specified ID is found, CreateTeam will return 0 and sql.ErrNoRows.
func (d *DB) CreateUser(name string, role int, teamID int) (id int64, err error) {
	var exists int
	err = d.db.QueryRow("SELECT 1 FROM teams WHERE id = ?;", teamID).Scan(&exists)
	if err != nil {
		return 0, err
	}

	res, err := d.db.Exec("INSERT INTO users(name, role, team_id) VALUES(?, ?, ?)", name, role, teamID)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
