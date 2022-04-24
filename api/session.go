package api

import (
	"net/http"
)

type newSessionResponse struct {
	ID int64 `json:"id"`
}

func (c Context) newSession(w http.ResponseWriter, _ *http.Request) {
	/*
		res, err := c.DB.Exec("INSERT INTO sessions DEFAULT VALUES;")
		if err != nil {
			log.Println(err)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		resp, err := json.Marshal(newSessionResponse{ID: id})
		if err != nil {
			log.Println(err)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Write(resp)
	*/
}
