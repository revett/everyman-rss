// Package rss provides HTTP handlers for RSS XML endpoints which Vercel will
// convert to serverless functions using the Go runtime.
// Note this cannot be within doc.go as Vercel sees that file as an endpoint.
package rss

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/feeds"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/revett/everyman-rss/internal/service"
	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/rs/zerolog/log"
)

const (
	cacheControl     = "s-maxage=300"
	cinemaQueryParam = "cinema"
)

// Films serves an RSS XML feed of the latest film releases from Everyman
// Cinema.
func Films(w http.ResponseWriter, r *http.Request) { // nolint:varnamelen
	log.Logger = commonLog.New()

	e := echo.New() // nolint:varnamelen
	e.Use(commonMiddleware.LoggerUsingZerolog(log.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.GET("/*", filmsHandler)
	e.ServeHTTP(w, r)
}

func filmsHandler(ctx echo.Context) error {
	cinemaSlug := ctx.QueryParam(cinemaQueryParam)
	if cinemaSlug == "" {
		m := fmt.Sprintf("request must have '%s' query param", cinemaQueryParam)
		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	cinemasClient, err := everyman.NewClientWithResponses(everyman.BaseWebURL)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to create everyman api client",
		)
	}

	cinemas, err := cinemasClient.CinemasWithResponse(context.TODO())
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"unable to request cinemas from everyman cinema api",
		)
	}

	var cinema everyman.Cinema

	for _, c := range *cinemas.JSON200 {
		if c.Slug() == cinemaSlug {
			cinema = c
			break
		}
	}

	if cinema.CinemaId == 0 {
		m := fmt.Sprintf("cinema '%s' does not exist", cinemaSlug)
		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	filmsClient, err := everyman.NewClientWithResponses(everyman.BaseAPIURL)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to create everyman api client",
		)
	}

	films, err := filmsClient.FilmsWithResponse(context.TODO(), cinema.CinemaId)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"unable to request films from everyman cinema api",
		)
	}

	feed, err := generateFeed(cinema, *films.JSON200)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to generate rss feed",
		)
	}

	ctx.Response().Header().Set("Cache-Control", cacheControl)
	return ctx.Blob(http.StatusOK, "application/xml", []byte(feed))
}

func generateFeed(cinema everyman.Cinema, films []everyman.Film) (string, error) {
	feed := feeds.Feed{
		Title: fmt.Sprintf("Everyman Cinema %s - Films", cinema.CinemaName),
		Description: fmt.Sprintf(
			"Latest film releases for Everyman Cinema %s.", cinema.CinemaName,
		),
		Link: &feeds.Link{
			Href: cinema.URL(),
		},
	}

	for _, film := range films {
		feed.Add(service.ConvertEverymanFilmToFeedItem(film))
	}

	rss, err := feed.ToRss()
	if err != nil {
		return "", fmt.Errorf("failed to generate rss feed: %w", err) // nolint:wrapcheck
	}

	return rss, nil
}
