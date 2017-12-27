package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func main() {
	commonHandlers := alice.New(loggingHandler)

	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	http.ListenAndServe(":8080", nil)
}


func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome /home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to /about")
}


