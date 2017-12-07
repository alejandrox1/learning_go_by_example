package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)


func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.Retrieve(id)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		return
	}
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return
	}
	err = post.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.Retrieve(id)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		return
	}
	err = post.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.Retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: Db}))
	server.ListenAndServe()
}
