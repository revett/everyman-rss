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

const (
	// See: https://vercel.com/docs/concepts/edge-network/caching#stale-while-revalidate
	cacheControl     = "s-maxage=300, stale-while-revalidate=3600"
	cinemaQueryParam = "cinema"
)

// Films serves an RSS XML feed of the latest film releases from Everyman
// Cinema.
func Films(w http.ResponseWriter, r *http.Request) {
	api.CommonMiddleware(films).ServeHTTP(w, r)
}

func films(w http.ResponseWriter, r *http.Request) { // nolint:funlen
	cinemaSlug := r.URL.Query().Get(cinemaQueryParam)

	if cinemaSlug == "" {
		err := fmt.Errorf("request must have '%s' query param", cinemaQueryParam)
		api.BadRequest(
			w, err, err.Error(),
		)

		return
	}

	cinemaClient, err := everyman.NewClientWithResponses(everyman.BaseWebURL)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to create everyman api client",
		)

		return
	}

	cinemas, err := cinemaClient.CinemasWithResponse(context.TODO())
	if err != nil {
		api.InternalServerError(
			w, err, "unable to request cinemas from everyman cinema api",
		)

		return
	}

	var cinemaID int

	for _, cinema := range *cinemas.JSON200 {
		if cinema.Slug() == cinemaSlug {
			cinemaID = cinema.CinemaId

			break
		}
	}

	if cinemaID == 0 {
		err := fmt.Errorf("cinema '%s' does not exist", cinemaSlug)
		api.BadRequest(
			w, err, err.Error(),
		)

		return
	}

	c, err := everyman.NewClientWithResponses(everyman.BaseAPIURL)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to create everyman api client",
		)

		return
	}

	films, err := c.FilmsWithResponse(context.TODO(), cinemaID)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to request films from everyman cinema api",
		)

		return
	}

	feed, err := generateFeed(*films.JSON200)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to generate rss feed",
		)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Cache-Control", cacheControl)
	w.Header().Set("Content-Type", "application/rss+xml")

	fmt.Fprint(w, feed)
}

func generateFeed(films []everyman.Film) (string, error) {
	f := feeds.Feed{
		Title:       "Everyman Cinema - Films",
		Description: "Latest film releases for Everyman Cinema",
		Link: &feeds.Link{
			Href: "https://www.everymancinema.com/film-listings",
		},
	}

	for _, film := range films {
		f.Add(
			service.ConvertEverymanFilmToFeedItem(film),
		)
	}

	rss, err := f.ToRss()
	if err != nil {
		return "", fmt.Errorf("failed to generate rss feed: %w", err)
	}

	return rss, nil
}
