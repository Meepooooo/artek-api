package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (c Context) getTeamData(w http.ResponseWriter, r *http.Request) {
	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Role string `json:"role"`
	}

	teamID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Team ID not specified or is not a number", http.StatusUnprocessableEntity)
		return
	}

	rows, err := c.DB.Query("SELECT id, name, role FROM users WHERE team_id = ?", teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
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
