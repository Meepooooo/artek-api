package main

import (
	"log"
	"net/http"

	"github.com/TaeKwonZeus/artek-api/api"
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/TaeKwonZeus/artek-api/db"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := db.Database(config.DB)
	if err != nil {
		log.Fatalln(err)
	}

	r := api.Router(db, config)

	http.ListenAndServe(":3000", r)
}
