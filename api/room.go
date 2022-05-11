package api

import (
	"encoding/json"
	"net/http"
)

func (c Context) createRoom(w http.ResponseWriter, r *http.Request) {
	res, err := c.DB.Exec("INSERT INTO ROOMS(date) VALUES(DATETIME('now'));")
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
		ID int `json:"id"`
	}{ID: int(id)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}
