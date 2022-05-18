package api

import (
	"fmt"
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

	r.Get("/", context.test)
	r.Route("/rooms", func(r chi.Router) {
		r.Post("/new", context.createRoom)
	})
	r.Route("/teams", func(r chi.Router) {
		r.Post("/new", context.createTeam)
	})
	r.Route("/users", func(r chi.Router) {
		r.Get("/", context.listUsers)
		r.Post("/new", context.createUser)
	})

	return r
}

func (c Context) test(w http.ResponseWriter, r *http.Request) {
	if c.DB == nil {
		http.Error(w, "Database connection doesn't exist", http.StatusInternalServerError)
		return
	}

	err := c.DB.Ping()
	if err != nil {
		http.Error(w, "Database failed to respond", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "This is a test endpoint. Use it to verify that the server and the database are online and working.")
}
