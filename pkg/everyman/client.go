package everyman

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the interface which defines the methods available for interacting
// with the Everyman Cinema API.
type Client interface {
	Cinemas() ([]Cinema, error)
	Films() ([]Film, error)
	FilmsByCinema(cinemaID int) ([]Film, error)
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Client.
func NewClient() Client {
	return client{
		baseURL:    BaseAPIURL,
		httpClient: http.DefaultClient,
	}
}

// Films implements the Client.Cinemas interface.
func (c client) Cinemas() ([]Cinema, error) {
	url := fmt.Sprintf("%s/%s", BaseWebURL, "cinemas")

	cinemas := []Cinema{}
	err := c.execute(url, &cinemas)
	if err != nil {
		return nil, err
	}

	return cinemas, nil
}

// Films implements the Client.Films interface.
func (c client) Films() ([]Film, error) {
	// TODO: this is limiting film search to a single cinema
	return c.FilmsByCinema(7925)
}

// Films implements the Client.FilmsByCinema interface.
func (c client) FilmsByCinema(cinemaID int) ([]Film, error) {
	url := fmt.Sprintf("%s/movies/13/%d", c.baseURL, cinemaID)

	films := []Film{}
	err := c.execute(url, &films)
	if err != nil {
		return nil, err
	}

	return films, nil
}

func (c client) execute(url string, v interface{}) error {
	r, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}
