package service_test

import (
	"testing"

	"github.com/gorilla/feeds"
	"github.com/revett/everyman-rss/internal/service"
	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/stretchr/testify/require"
)

func TestConvertEverymanFilmToFeedItem(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		film everyman.Film
		want *feeds.Item
	}{
		"Success": {
			film: everyman.Film{
				FilmId: 987,
				Title:  "Dune",
				MediaItems: everyman.FilmMediaItems{
					QuadStill: "http://images.mymovies.net/images/film/cin/stills/531x329/fid20292/1.jpg",
				},
				Synopsis:     "<p>Oscar nominee Denis Villeneuve...</p>",
				FriendlyName: "dune",
			},
			want: &feeds.Item{
				Id:          "987",
				Title:       "Dune",
				Description: `<img src="http://images.mymovies.net/images/film/cin/stills/531x329/fid20292/1.jpg" /><br><br><p>Oscar nominee Denis Villeneuve...</p>`, // nolint:lll
				Link: &feeds.Link{
					Href: "https://everymancinema.com/film-info/dune",
				},
			},
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			item := service.ConvertEverymanFilmToFeedItem(tc.film)
			require.Equal(t, tc.want, item)
		})
	}
}
