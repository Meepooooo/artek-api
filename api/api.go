package api

import (
	"fmt"
	"net/http"

	"github.com/TaeKwonZeus/artek-api/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Env struct {
	DB *db.DB
}

func Router(e Env) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/", e.test)
	r.Route("/rooms", func(r chi.Router) {
		r.Post("/new", e.createRoom)
	})
	r.Route("/teams", func(r chi.Router) {
		r.Get("/{id}", e.getTeam)
		r.Post("/new", e.createTeam)
	})
	r.Route("/users", func(r chi.Router) {
		r.Post("/new", e.createUser)
	})

	return r
}

func (e Env) test(w http.ResponseWriter, r *http.Request) {
	if e.DB == nil {
		http.Error(w, "Database connection doesn't exist", http.StatusInternalServerError)
		return
	}

	err := e.DB.Ping()
	if err != nil {
		http.Error(w, "Database failed to respond", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "This is a test endpoint. Use it to verify that the server and the database are online and working.")
}
