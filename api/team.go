package api

import (
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
