package main

import (
	"html/template"
	"log"
	"net/http"
)


func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
