package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (c Context) createUser(w http.ResponseWriter, r *http.Request) {
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
	err = c.DB.QueryRow("SELECT 1 FROM teams WHERE id = ?;", body.TeamID).Scan(&exists)
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

	res, err := c.DB.Exec("INSERT INTO users(name, role, team_id) VALUES(?, ?, ?)", body.Name, body.Role, body.TeamID)
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
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Role int    `json:"role"`
	}{ID: id, Name: body.Name, Role: body.Role}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
