package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/justinas/alice"

	"github.com/gorilla/handlers"
)

func loggingHandler(next http.Handler) http.Handler {
	logFile, err := os.OpenFile("server1.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("Errro write file ligFile: %v", err)
	}
	return handlers.LoggingHandler(logFile, next)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
        <html>
            <head>
                <title>Index</title>
            </head>
            <body>
                Hello Gopher!
            </body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
        <html>
            <head>
                <title>About</title>
            </head>
            <body>
                Go Web development with HTTP Middleware
            </body>
</html>`)
}

func iconHandler1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler1)
	indexHandler1 := http.HandlerFunc(index)
	aboutHandler1 := http.HandlerFunc(about)

	commonHandlers := alice.New(loggingHandler, handlers.CompressHandler)

	http.Handle("/", commonHandlers.ThenFunc(indexHandler1))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler1))

	server := &http.Server{
		Addr: ":9003",
	}
	log.Println("Running on http://localhost" + server.Addr)
	server.ListenAndServe()
}
