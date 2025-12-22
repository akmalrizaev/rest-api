package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
		return
	case http.MethodPost:

	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Students Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Students Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Students Route"))
		return
	}

	w.Write([]byte("Hello from Students Route"))
}
