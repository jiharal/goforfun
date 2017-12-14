package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/gorilla/handlers"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute index handler")
	fmt.Fprintf(w, "Welcome !")
}

func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute about handler")
	fmt.Fprintf(w, "Go Midleware !")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error create log file : %v", err)
	}

	http.Handle("/", handlers.LoggingHandler(logFile, handlers.CompressHandler(indexHandler)))
	http.Handle("/about", handlers.LoggingHandler(logFile, handlers.CompressHandler(aboutHandler)))

	server := &http.Server{
		Addr: ":9002",
	}
	log.Println("Running on http://localhost" + server.Addr)
	server.ListenAndServe()
}
