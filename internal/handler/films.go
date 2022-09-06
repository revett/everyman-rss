package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/revett/everyman-rss/internal/service"
	"github.com/revett/everyman-rss/pkg/everyman"
)

const (
	cacheControl     = "s-maxage=300"
	cinemaQueryParam = "cinema"
)

// Films serves an RSS XML feed of the latest film releases from Everyman
// Cinema.
func Films(ctx echo.Context) error { //nolint:cyclop,funlen
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

	feed, err := service.GenerateFeed(cinema, *films.JSON200)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to generate rss feed",
		)
	}

	ctx.Response().Header().Set("Cache-Control", cacheControl)

	if err := ctx.Blob(http.StatusOK, "application/xml", []byte(feed)); err != nil {
		return fmt.Errorf(
			"failed to send xml blob content with status code: %w", err,
		)
	}

	return nil
}
