package service_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/revett/everyman-rss/internal/service"
	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/stretchr/testify/require"
)

func TestGenerateFeed(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cinema everyman.Cinema
		films  []everyman.Film
		err    require.ErrorAssertionFunc
	}{
		"Success": {
			cinema: everyman.Cinema{
				CinemaName: "Winchester",
			},
			films: []everyman.Film{
				{
					FilmId: 39878,
					Title:  "Dune",
					MediaItems: everyman.FilmMediaItems{
						QuadStill: "http://images.mymovies.net/images/film/cin/stills/531x329/fid20292/1.jpg",
					},
					Synopsis:     "<p>Oscar nominee Denis Villeneuve...</p>",
					FriendlyName: "dune",
				},
				{
					FilmId: 39309,
					Title:  "No Time To Die",
					MediaItems: everyman.FilmMediaItems{
						QuadStill: "https://legacylivefilmdbstorage.blob.core.windows.net/evm-circuit-13/4-39309-07352071-74b4-4c2a-876a-644e54733913.jpg", // nolint:lll
					},
					Synopsis:     "<p>Bond has left active service...</p>",
					FriendlyName: "no-time-to-die",
				},
			},
			err: require.NoError,
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			b, err := os.ReadFile(
				filepath.Join("testdata", t.Name()+".golden"),
			)
			require.NoError(t, err)
			want := strings.TrimSuffix(string(b), "\n")

			feed, err := service.GenerateFeed(tc.cinema, tc.films)
			tc.err(t, err)
			require.Equal(t, want, feed)
		})
	}
}
