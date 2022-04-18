package api

import (
	"net/http"

	"github.com/TaeKwonZeus/artek-api/api/handlers"
	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB, config config.Config) http.Handler {
	handler := handlers.Handler{DB: db}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", handler.Test)
	})

	return r
}
