package data

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS teams(
    id INT PRIMARY KEY,
	name TEXT,
	room_id INT,
	FOREIGN KEY(room_id) REFERENCES rooms(id)
);
`

func Database(config config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", config.DSN)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
