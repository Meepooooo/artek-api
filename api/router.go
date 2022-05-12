package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(context Context) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/", test)
	r.Route("/rooms", func(r chi.Router) {
		r.Post("/new", context.createRoom)
	})
	r.Route("/teams", func(r chi.Router) {
		r.Post("/new", context.createTeam)
	})
	r.Route("/users", func(r chi.Router) {
		r.Post("/new", context.createUser)
	})

	return r
}
