package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public", http.StripPrefix("/public", fs))
}

// How to write http midleware
func midlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello ini mid 1")
		next.ServeHTTP(w, r)
		fmt.Println("Hello ini mid 2")

		fmt.Println("Apa gunanya midleware ?")
	})
}
