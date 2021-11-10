package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	handler "github.com/revett/everyman-rss/api"
)

const (
	addr    = "127.0.0.1:5691"
	timeout = 5
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index)
	r.HandleFunc("/films", handler.Films)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: timeout * time.Second,
		ReadTimeout:  timeout * time.Second,
	}

	log.Println(addr)

	log.Fatal(
		srv.ListenAndServe(),
	)
}
