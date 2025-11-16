package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming orders ")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users")
	})

	const port string = ":3000"

	fmt.Println("Server Listening on port: ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}

}
