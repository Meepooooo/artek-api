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

	var exists int
	err = e.DB.QueryRow("SELECT 1 FROM teams WHERE id = ?;", body.TeamID).Scan(&exists)
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

	res, err := e.DB.Exec("INSERT INTO users(name, role, team_id) VALUES(?, ?, ?)", body.Name, body.Role, body.TeamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID int64 `json:"id"`
	}{ID: id}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
