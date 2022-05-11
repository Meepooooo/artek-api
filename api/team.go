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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Handle all internal server errors
	resp, err := func() ([]byte, error) {
		res, err := c.DB.Exec("INSERT INTO teams(name, room_id) VALUES(?, ?);", body.Name, body.RoomId)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		resp, err := json.Marshal(struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}{ID: id, Name: body.Name})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}
