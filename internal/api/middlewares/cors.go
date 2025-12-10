package middlewares

import (
	"fmt"
	"net/http"
)

func Cors(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Accept-Encoding")
		fmt.Println(origin)

		next.ServeHTTP(w, r)
	})

}
