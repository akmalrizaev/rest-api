package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Router")
		w.Write([]byte("Hello Root Router"))
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET Method on Teachers Route"))
			return
		case http.MethodPost:

			// Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			fmt.Println("Form", r.Form)

			// Prepare response data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				// response[key] = value
				response[key] = value[0]

			}

			fmt.Println("Processed Response Map:", response)

			w.Write([]byte("Hello POST Method on Teachers Route"))
			return
		case http.MethodPut:
			w.Write([]byte("Hello PUT Method on Teachers Route"))
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH Method on Teachers Route"))
			return
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE Method on Teachers Route"))
			return
		}

		if r.Method == http.MethodGet {
			w.Write([]byte("Hello GET Method on Teachers Route"))
			return
		}
		w.Write([]byte("Hello from Teachers Route"))
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Students Route"))
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Execs Route"))
	})

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
