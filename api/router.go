package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(context Context) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/", context.test)
		r.Route("/session", func(r chi.Router) {
			r.Post("/new", context.newSession)
		})
	})

	return r
}
