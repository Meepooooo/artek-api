package handlers

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	DB     *sqlx.DB
	Config config.Config
}
