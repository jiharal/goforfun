package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Midleware

func loggingHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started, %s, %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed, %s in %v", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute index handler")
	fmt.Fprintf(w, "Happy learning.")
}

func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute about handler")
	fmt.Fprintf(w, "Semua tentang middleware")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)

	http.Handle("/", loggingHandle(indexHandler))
	http.Handle("/about", loggingHandle(aboutHandler))

	server := &http.Server{
		Addr: ":9000",
	}

	log.Println("Listening ...", "http://localhost"+server.Addr)
	server.ListenAndServe()
}
