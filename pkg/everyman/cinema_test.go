package everyman_test

import (
	"testing"

	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/stretchr/testify/require"
)

func TestCinemaSlug(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cinema everyman.Cinema
		want   string
	}{
		"Simple": {
			cinema: everyman.Cinema{
				CinemaName: "Winchester",
			},
			want: "winchester",
		},
		"TwoWords": {
			cinema: everyman.Cinema{
				CinemaName: "Maida Vale",
			},
			want: "maida-vale",
		},
		"Punctuation": {
			cinema: everyman.Cinema{
				CinemaName: "Manchester St. John's",
			},
			want: "manchester-st-johns",
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			slug := tc.cinema.Slug()
			require.Equal(t, tc.want, slug)
		})
	}
}

func TestCinemaURL(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cinema everyman.Cinema
		want   string
	}{
		"Simple": {
			cinema: everyman.Cinema{
				CinemaName: "Winchester",
			},
			want: "https://everymancinema.com/winchester",
		},
		"TwoWords": {
			cinema: everyman.Cinema{
				CinemaName: "Maida Vale",
			},
			want: "https://everymancinema.com/maida-vale",
		},
		"Punctuation": {
			cinema: everyman.Cinema{
				CinemaName: "Manchester St. John's",
			},
			want: "https://everymancinema.com/manchester-st-johns",
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			url := tc.cinema.URL()
			require.Equal(t, tc.want, url)
		})
	}
}
