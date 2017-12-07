package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// Struct for html templates.
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.Int("addr", 8080, "Connection port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.Handle("/", &templateHandler{filename: "chat.html"})

	r := newRoom()
	mux.Handle("/room", r)
	go r.run()

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *addr),
		Handler: mux,
	}
	log.Println("Listening to port ", *addr)
	log.Fatal(server.ListenAndServe())
}
