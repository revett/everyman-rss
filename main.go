package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
	"github.com/gorilla/mux"
	"github.com/revett/everyman-rss/pkg/everyman"
)

const (
	addr    = "127.0.0.1:5691"
	timeout = 5
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/films", filmsHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: timeout * time.Second,
		ReadTimeout:  timeout * time.Second,
	}

	log.Fatal(
		srv.ListenAndServe(),
	)
}

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	f := feeds.Feed{
		Title:       "Everyman Cinema - Films",
		Description: "Latest film releases for Everyman Cinema",
		Link: &feeds.Link{
			Href: "https://www.everymancinema.com/film-listings",
		},
	}

	c := everyman.NewClient()

	films, err := c.Films()
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"unable to request films from everyman cinema api",
			http.StatusInternalServerError,
		)
	}

	for _, film := range films {
		f.Add(
			&feeds.Item{
				Id:          strconv.Itoa(film.ID),
				Title:       film.Title,
				Description: film.Teaser,
				Link: &feeds.Link{
					Href: film.URL(),
				},
			},
		)
	}

	rss, err := f.ToRss()
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to generate rss feed", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, rss)
}
