package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	commonLog "github.com/revett/common/log"
	handler "github.com/revett/everyman-rss/api"
	"github.com/rs/zerolog/log"
)

const (
	addr    = "127.0.0.1:5691"
	timeout = 5
)

func main() {
	log.Logger = commonLog.New()
	log.Info().Msg(addr)

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

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
