package api

import (
	"encoding/json"
	"net/http"
)

func (c Context) createTeam(w http.ResponseWriter, _ *http.Request) {
	res, err := c.DB.Exec("INSERT INTO teams DEFAULT VALUES")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := json.Marshal(struct {
		ID int `json:"id"`
	}{ID: int(id)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}
