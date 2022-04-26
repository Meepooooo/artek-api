package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(context Context) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", test)
	r.Route("/room", func(r chi.Router) {
		r.Post("/new", context.createRoom)
	})
	r.Route("/team", func(r chi.Router) {
		r.Post("/new", context.createTeam)
	})

	return r
}
