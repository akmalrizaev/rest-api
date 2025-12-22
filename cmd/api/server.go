package main

import (
	"fmt"
	"log"
	"net/http"
	"simpleapi/internal/api/middlewares"
	"simpleapi/internal/api/router"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {

	port := ":3000"

	mux := router.Router()

	// rl := middlewares.NewRateLimiter(5, time.Minute)

	fmt.Println("Server is running on port", port)
	// err := http.ListenAndServe(port, rl.Middleware(middlewares.Compression(middlewares.ResponseTimeMiddleware(middlewares.Cors(middlewares.SecurityHeaders(mux))))))
	// err := http.ListenAndServe(port, ApplyMiddlewares(mux, middlewares.SecurityHeaders, middlewares.Cors, middlewares.ResponseTimeMiddleware, middlewares.Compression, rl.Middleware))
	err := http.ListenAndServe(port, middlewares.SecurityHeaders(mux))
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
