package api

import "net/http"

func (Context) test(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World!"))
}
