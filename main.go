package main

import (
	"net/http"
	"sync"
	"time"
	"gopkg.in/tylerb/graceful.v1"
	"log"
	"github.com/edustor/gen/source"
	"github.com/urfave/negroni"
)

type handler struct {
	w sync.WaitGroup
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	graceful.Run(":8080", 1 * time.Second, n)
}

func handleRoot(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Processing request %v from %v", req.URL, req.RemoteAddr)
	err := source.GenPdf(rw)
	if (err != nil) {
		log.Panic(err)
	}
}
