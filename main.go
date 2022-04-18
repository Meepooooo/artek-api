package main

import (
	"log"
	"net/http"

	"github.com/TaeKwonZeus/artek-api/config"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database(config.DB)
	if err != nil {
		log.Fatalln(err)
	}

	r := router(db, config)

	http.ListenAndServe(":3000", r)
}
