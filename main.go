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

	db, err := db.Database(config.DBLocation)
	if err != nil {
		log.Fatalln(err)
	}

	context := api.Context{DB: db}
	r := api.Router(context)

	log.Printf("available at port %d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
