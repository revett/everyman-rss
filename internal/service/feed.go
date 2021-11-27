package service

import (
	"fmt"

	"github.com/gorilla/feeds"
	"github.com/revett/everyman-rss/pkg/everyman"
)

// GenerateFeed generates a RSS feed, containing films.
func GenerateFeed(cinema everyman.Cinema, films []everyman.Film) (string, error) {
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
		feed.Add(ConvertEverymanFilmToFeedItem(film))
	}

	rss, err := feed.ToRss()
	if err != nil {
		return "", fmt.Errorf("failed to generate rss feed: %w", err) // nolint:wrapcheck
	}

	return rss, nil
}
