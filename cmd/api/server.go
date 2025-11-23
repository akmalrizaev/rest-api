package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

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

			/*

				Parse form data (necessary for x-www-form-urlencoded)
				err := r.ParseForm()
				if err != nil {
					http.Error(w, "Error parsing form", http.StatusBadRequest)
					return
				}

				fmt.Println("Form", r.Form)

				Prepare response data
				response := make(map[string]interface{})
				for key, value := range r.Form {
					response[key] = value
					response[key] = value[0]

				}

				fmt.Println("Processed Response Map:", response)

				RAW Body
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return
				}
				defer r.Body.Close()

				fmt.Println("RAW Body without converting to string:", body)
				fmt.Println("RAW Body", string(body))

				if we expect json data, then unmarshal it
				var userInstance User
				err = json.Unmarshal(body, &userInstance)
				if err != nil {
					return
				}

				fmt.Println("JSON Unmarshaling:", userInstance)
				fmt.Println("Received user name as:", userInstance.Name)

				Access the request details
				fmt.Println("Access the request details:")
				fmt.Println("Body:", r.Body)
				fmt.Println("Form", r.Form)
				fmt.Println("Header", r.Header)
				fmt.Println("Context", r.Context())
				fmt.Println("ContentLength", r.ContentLength)
				fmt.Println("Host", r.Host)
				fmt.Println("Method", r.Method)
				fmt.Println("Protocol", r.Proto)
				fmt.Println("Remote Addr", r.RemoteAddr)
				fmt.Println("Request URI", r.RequestURI)
				fmt.Println("TLS", r.TLS)
				fmt.Println("Trailers", r.Trailer)
				fmt.Println("Transfer Encoding", r.TransferEncoding)
				fmt.Println("URL", r.URL)
				fmt.Println("User Agent", r.UserAgent())
				fmt.Println("Port", r.URL.Port())
				fmt.Println("URL", r.URL.Scheme)

				w.Write([]byte("Hello POST Method on Teachers Route"))
				return

			*/

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
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET Method on Execs Route"))
			return
		case http.MethodPost:

		case http.MethodPut:
			w.Write([]byte("Hello PUT Method on Execs Route"))
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH Method on Execs Route"))
			return
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE Method on Execs Route"))
			return
		}

		w.Write([]byte("Hello from Execs Route"))
	})

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
