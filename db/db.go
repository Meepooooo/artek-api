package db

import (
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

func Database(user string, password string, dbName string) (*sqlx.DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbName)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile("db/start.sql")
	if err != nil {
		return nil, err
	}

	db.MustExec(string(bytes))

	return db, nil
}
