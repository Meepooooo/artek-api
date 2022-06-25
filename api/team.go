package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (e Env) getTeam(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Team ID not specified or is not a number", http.StatusUnprocessableEntity)
		return
	}

	team, err := e.DB.GetTeam(id)
	switch err {
	case sql.ErrNoRows:
		http.Error(w, "Team with given ID does not exist", http.StatusUnprocessableEntity)
		return
	case nil:
		break
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}

func (e Env) createTeam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name   string
		RoomID int
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	id, err := e.DB.CreateTeam(body.Name, body.RoomID)
	switch err {
	case sql.ErrNoRows:
		http.Error(w, "Room with given ID does not exist", http.StatusUnprocessableEntity)
		return
	case nil:
		break
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID int64 `json:"id"`
	}{ID: id}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (e Env) setTeamBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Team ID not specified or is not a number", http.StatusUnprocessableEntity)
		return
	}

	balance, err := e.DB.SetTeamBalance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}
