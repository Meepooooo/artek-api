package handlers

import "github.com/jmoiron/sqlx"

type Handler struct {
	DB *sqlx.DB
}
