package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (c Context) getTeam(w http.ResponseWriter, r *http.Request) {
	type team struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		RoomID string `json:"roomId"`
	}

	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Role string `json:"role"`
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Team ID not specified or is not a number", http.StatusUnprocessableEntity)
		return
	}

	var row team
	err = c.DB.QueryRow("SELECT id, name, room_id FROM teams WHERE id = ?;", id).Scan(&row.ID, &row.Name, &row.RoomID)
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

	rows, err := c.DB.Query("SELECT id, name, role FROM users WHERE team_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	resp := struct {
		team
		Users []user `json:"users"`
	}{row, users}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (c Context) createTeam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name   string
		RoomID int
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var exists int
	err = c.DB.QueryRow("SELECT 1 FROM rooms WHERE id = ?;", body.RoomID).Scan(&exists)
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

	res, err := c.DB.Exec("INSERT INTO teams(name, room_id) VALUES(?, ?);", body.Name, body.RoomID)
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
	}{ID: id, Name: body.Name}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
