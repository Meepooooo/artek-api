package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (e Env) createUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name   string
		Role   int
		TeamID int
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := e.DB.CreateUser(body.Name, body.Role, body.TeamID)
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

	resp := struct {
		ID int64 `json:"id"`
	}{ID: id}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
