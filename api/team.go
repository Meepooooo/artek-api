package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (c Context) createTeam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name   string
		RoomId int
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var exists int
	err = c.DB.QueryRow("SELECT 1 FROM rooms WHERE id = ?;", body.RoomId).Scan(&exists)
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

	res, err := c.DB.Exec("INSERT INTO teams(name, room_id) VALUES(?, ?);", body.Name, body.RoomId)
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
	json.NewEncoder(w).Encode(&resp)
}
