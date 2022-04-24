package api

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/TaeKwonZeus/artek-api/data"
	"github.com/dgraph-io/badger/v3"
)

type Context struct {
	DB     *badger.DB
	Config config.APIConfig
	State  *data.State
}
