package main

import (
	"fmt"
	"log"
	"net/http"
)

func midlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MidlewareFirst, Before Handler")
		next.ServeHTTP(w, r)
		log.Println("MidlewareFirst, After Handler")
	})
}

func midlerwareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("midlewareSecond - Before Handler")

		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "pass123" {
				log.Println("Authorized to the system")
				next.ServeHTTP(w, r)
			} else {
				log.Println("Failure to authorize to the system")
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
		log.Println("midlewaresecond - After Handler")
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute index handler")
	fmt.Fprintf(w, "Ini halaman index")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute message handler")
	fmt.Fprintf(w, "HTTP midleware is awesome")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	http.Handle("/", midlewareFirst(midlerwareSecond(http.HandlerFunc(index))))
	http.Handle("/message", midlewareFirst(midlerwareSecond(http.HandlerFunc(message))))

	server := &http.Server{
		Addr: ":9090",
	}
	log.Println("Listener: ", "http://localhost"+server.Addr)
	server.ListenAndServe()
}
