package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/feeds"
	"github.com/revett/everyman-rss/internal/api"
	"github.com/revett/everyman-rss/internal/service"
	"github.com/revett/everyman-rss/pkg/everyman"
)

// See: https://vercel.com/docs/concepts/edge-network/caching#stale-while-revalidate
const cacheControl = "s-maxage=300, stale-while-revalidate=3600"

// Films serves an RSS XML feed of the latest film releases from Everyman
// Cinema.
func Films(w http.ResponseWriter, r *http.Request) {
	api.CommonMiddleware(films).ServeHTTP(w, r)
}

func films(w http.ResponseWriter, r *http.Request) {
	f := feeds.Feed{
		Title:       "Everyman Cinema - Films",
		Description: "Latest film releases for Everyman Cinema",
		Link: &feeds.Link{
			Href: "https://www.everymancinema.com/film-listings",
		},
	}

	c, err := everyman.NewClientWithResponses(everyman.BaseAPIURL)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to create everyman api client",
		)
		return
	}

	films, err := c.FilmsWithResponse(context.TODO(), 7925)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to request films from everyman cinema api",
		)
		return
	}

	for _, film := range *films.JSON200 {
		f.Add(
			service.ConvertEverymanFilmToFeedItem(film),
		)
	}

	rss, err := f.ToRss()
	if err != nil {
		api.InternalServerError(
			w, err, "unable to generate rss feed",
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Cache-Control", cacheControl)
	w.Header().Set("Content-Type", "application/rss+xml")

	fmt.Fprint(w, rss)
}
