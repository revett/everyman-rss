package everyman_test

import (
	"testing"

	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/stretchr/testify/require"
)

func TestFilmURL(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		film everyman.Film
		want string
	}{
		"Success": {
			film: everyman.Film{
				FriendlyName: "dune",
			},
			want: "https://everymancinema.com/film-info/dune",
		},
	}

	for n, testCase := range tests {
		tc := testCase

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			url := tc.film.URL()
			require.Equal(t, tc.want, url)
		})
	}
}
