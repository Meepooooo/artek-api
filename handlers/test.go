package handlers

import "net/http"

func (h Context) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
