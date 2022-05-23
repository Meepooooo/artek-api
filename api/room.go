package api

import (
	"encoding/json"
	"net/http"
)

func (e Env) createRoom(w http.ResponseWriter, r *http.Request) {
	id, err := e.DB.CreateRoom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID int `json:"id"`
	}{ID: int(id)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
