package api

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/TaeKwonZeus/artek-api/data"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	DB     *sqlx.DB
	Config config.APIConfig
	State  *data.State
}
