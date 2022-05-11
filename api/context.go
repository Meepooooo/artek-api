package api

import "database/sql"

type Context struct {
	DB *sql.DB
}
