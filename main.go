package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TaeKwonZeus/artek-api/api"
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/TaeKwonZeus/artek-api/db"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := db.Open(config.DB)
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	env := api.Env{DB: db}
	r := api.Router(env)

	log.Printf("available at port %d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
