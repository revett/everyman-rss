package service

import (
	"fmt"
	"strconv"

	"github.com/gorilla/feeds"
	"github.com/revett/everyman-rss/pkg/everyman"
)

// ConvertEverymanFilmToFeedItem converts an everyman.Film type to an RSS
// feeds.Item type.
func ConvertEverymanFilmToFeedItem(film everyman.Film) *feeds.Item {
	return &feeds.Item{
		Id:    strconv.Itoa(film.FilmId),
		Title: film.Title,
		Description: fmt.Sprintf(
			`<img src="%s" /><br><br>%s`, film.MediaItems.QuadStill, film.Synopsis,
		),
		Link: &feeds.Link{
			Href: film.URL(),
		},
	}
}
