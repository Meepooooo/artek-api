package handlers

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB     *sqlx.DB
	Config config.Config
}
