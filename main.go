package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TaeKwonZeus/artek-api/api"
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/TaeKwonZeus/artek-api/data"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = data.Database(config.DB)
	if err != nil {
		log.Fatalln(err)
	}

	context := api.Context{Config: config.API}
	r := api.Router(context)

	log.Printf("available at port %d", config.API.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.API.Port), r)
}
