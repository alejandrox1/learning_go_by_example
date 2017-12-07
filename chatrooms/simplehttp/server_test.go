package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	//"os"
	"strings"
	"testing"
)

/*
var mux *http.ServeMux
var writer *httptest.ResponseRecorder
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
func setUp() {
	mux := http.NewServeMux()
	writer := httptest.NewRecorder()
	/* 
	jsonRequest := strings.NewReader(`{"content": "Hi", "author": "Jorge"}`)
	request, err := http.NewRequest("POST", "/post/1", jsonRequest)
	if err != nil {
		t.Error("Error sending post request: ", err)
	}
	
}
*/

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	err := json.NewDecoder(writer.Body).Decode(&post)
	if err != nil {
		t.Error("Error decoding JSON post: ", err)
	}
	if post.Id != 1 {
		t.Error("Error retrieving JSON post, got the wrong post id")
	}
}


func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	jsonRequest := strings.NewReader(`{"content": "Updated post", "author": "Jorge"}`)
	request, _ := http.NewRequest("PUT", "/post/1", jsonRequest)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
