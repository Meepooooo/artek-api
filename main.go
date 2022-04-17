package main

import (
	"log"
	"net/http"

	"github.com/TaeKwonZeus/artek-api/api"
	"github.com/TaeKwonZeus/artek-api/db"
	_ "github.com/lib/pq"
)

func main() {
	db, err := db.Database("postgres", "postgres", "artek_api")
	if err != nil {
		log.Fatalln(err)
	}

	r := api.Router(db)

	http.ListenAndServe(":3000", r)
}
