package main

import (
	"net/http"

	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/TaeKwonZeus/artek-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func router(db *sqlx.DB, config config.Config) http.Handler {
	handler := handlers.Context{DB: db, Config: config}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", handler.Test)
	})

	return r
}
