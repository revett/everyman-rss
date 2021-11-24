package handler

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
func Films(w http.ResponseWriter, r *http.Request) {
	log.Logger = commonLog.New()

	e := echo.New()
	e.Use(commonMiddleware.LoggerUsingZerolog(log.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.GET("/films", filmsHandler)
	e.ServeHTTP(w, r)
}

func filmsHandler(ctx echo.Context) error {
	cinemaSlug := ctx.QueryParam(cinemaQueryParam)
	if cinemaSlug == "" {
		m := fmt.Sprintf("request must have '%s' query param", cinemaQueryParam)
		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	cinemaClient, err := everyman.NewClientWithResponses(everyman.BaseWebURL)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to create everyman api client",
		)
	}

	cinemas, err := cinemaClient.CinemasWithResponse(context.TODO())
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"unable to request cinemas from everyman cinema api",
		)
	}

	var cinemaID int

	for _, cinema := range *cinemas.JSON200 {
		if cinema.Slug() == cinemaSlug {
			cinemaID = cinema.CinemaId
			break
		}
	}

	if cinemaID == 0 {
		m := fmt.Sprintf("cinema '%s' does not exist", cinemaSlug)
		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	c, err := everyman.NewClientWithResponses(everyman.BaseAPIURL)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to create everyman api client",
		)
	}

	films, err := c.FilmsWithResponse(context.TODO(), cinemaID)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"unable to request films from everyman cinema api",
		)
	}

	feed, err := generateFeed(*films.JSON200)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to generate rss feed",
		)
	}

	ctx.Response().Header().Set("Cache-Control", cacheControl)
	return ctx.Blob(http.StatusOK, "application/xml", []byte(feed))
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
