package db

const (
	roleMechanic = iota + 1
	roleDrunkard
	roleMage
)

var roles = map[int]role{
	roleMechanic: {
		defaults: Resources{2, 2, 2, 2},
		spending: Resources{},
		earning:  Resources{},
	},
	roleDrunkard: {
		defaults: Resources{3, 3, 3, 3},
	},
	roleMage: {
		defaults: Resources{1, 1, 1, 1},
	},
}

type role struct {
	defaults Resources
	spending Resources
	earning  Resources
}

// User represents a quest user.
// A user belongs to a team.
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Role   int    `json:"role"`
	TeamID int    `json:"teamId,omitempty"`
}

// CreateUser creates a new user with the specified name, role and team ID.
// If no team with the specified ID is found, CreateUser will return 0 and sql.ErrNoRows.
func (d *DB) CreateUser(name string, role int, teamID int) (id int64, err error) {
	err = d.db.QueryRow("SELECT 1 FROM teams WHERE id = ?;", teamID).Scan(new(byte))
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
