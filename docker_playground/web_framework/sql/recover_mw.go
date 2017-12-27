package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func main() {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)

	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	http.Handle("/admin", commonHandlers.Append(uthHandler).ThenFunc(adminHandler))
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

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v\n", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header().Get("Authorization")
		user, err := getUser(authToken)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		context.Set(r, "user", user)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}


func adminHandler(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	json.NewEncoder(w).Encode(user)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome /home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to /about")
}


