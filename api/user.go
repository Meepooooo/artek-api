package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func (c Context) listUsers(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}

	teamID, err := strconv.Atoi(r.URL.Query().Get("teamid"))
	if err != nil {
		http.Error(w, "Query parameter teamid not specified or is not a number", http.StatusUnprocessableEntity)
		return
	}

	rows, err := c.DB.Query("SELECT name, role FROM users WHERE team_id = ?", teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.Name, &user.Role); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		users = append(users, user)
	}

	resp := struct {
		TeamID int    `json:"teamId"`
		Users  []user `json:"users"`
	}{teamID, users}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

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
