package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)



func main() {
	http.HandleFunc("/", handler)
	http.Handle("/home", handlerstruct{})
	http.HandleFunc("/timer", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http.HandleFunc")
}

type handlerstruct struct {}

func (h handlerstruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http.Handle")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	fmt.Fprintf(w, "welcome")
	t2 := time.Now()

	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}


