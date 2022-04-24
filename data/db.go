package data

import (
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/dgraph-io/badger/v3"
)

func Database(config config.DBConfig) (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions(config.Location))
	if err != nil {
		return nil, err
	}

	return db, nil
}
