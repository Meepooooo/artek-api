package handlers

import "net/http"

func (Context) Test(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World!"))
}
