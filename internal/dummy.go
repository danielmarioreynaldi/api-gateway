package internal

import "net/http"

func Dummy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dummy gateway working"))
}
