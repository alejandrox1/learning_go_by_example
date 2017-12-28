// curl -Lv --proxy https://localhost:8080 --proxy-cacert server.pem https://google.com
package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

// Two TCP connections (client -> proxy and proxy -> destination server)
func handleTunneling(w http.ResponseWriter, r *http.Request) {
	log.Println("R.Host: ", r.Host)
	// Set the connection to a destination server.
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}


func copyHeader(dest, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dest.Add(k, v)
		}
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	// 
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}


func main() {
	var pemPath string
	var keyPath string
	var proto string
	flag.StringVar(&pemPath, "pem", "server.pem", "path to pem file")
	flag.StringVar(&keyPath, "key", "server.key", "path to key file")
	flag.StringVar(&proto, "proto", "https", "Proxy protocol (http || https)")
	flag.Parse()

	// Quit if wrong protocol
	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either HTTP or HTTPS")
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("r.Method: ", r.Method)
			if r.Method == http.MethodConnect {
				log.Println("handleTunneling is handling request")
				handleTunneling(w, r)
			} else {
				log.Println("handleHTTP is handling request")
				handleHTTP(w, r)
			}
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if proto == "http" {
		log.Fatal(server.ListenAndServe)
	} else if proto == "https" {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}
}
