package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Router")
	w.Write([]byte("Hello Root Router"))
}
