package main

import (
	"net/http"
	"sync"
	"time"
	"gopkg.in/tylerb/graceful.v1"
	"log"
)

type handler struct {
	w sync.WaitGroup
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	graceful.Run(":8080", 1*time.Second, mux)
}

func handleRoot(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Processing request %v from %v", req.URL, req.RemoteAddr)
	rw.Write([]byte("Hello world"))
}
