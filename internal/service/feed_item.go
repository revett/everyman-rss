package service

import (
	"fmt"
	"strconv"

	"github.com/gorilla/feeds"
	"github.com/revett/everyman-rss/pkg/everyman"
)

// ConvertEverymanFilmToFeedItem converts an everyman.Film type to an RSS
// feeds.Item type.
func ConvertEverymanFilmToFeedItem(f everyman.Film) *feeds.Item {
	return &feeds.Item{
		Id:    strconv.Itoa(f.FilmId),
		Title: f.Title,
		Description: fmt.Sprintf(
			`<img src="%s" /><br><br>%s`, f.MediaItems.QuadStill, f.Synopsis,
		),
		Link: &feeds.Link{
			Href: fmt.Sprintf("%s/film-info/%s", everyman.BaseWebURL, f.FriendlyName),
		},
	}
}
