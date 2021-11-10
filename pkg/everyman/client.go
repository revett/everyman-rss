package everyman

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the interface which defines the methods available for interacting
// with the API.
type Client interface {
	Films() ([]Film, error)
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

const baseAPIURL = "https://movieeverymanapi.peachdigital.com"

// NewClient creates a new Client.
func NewClient() Client {
	return client{
		baseURL:    baseAPIURL,
		httpClient: http.DefaultClient,
	}
}

// Films implements the Client.Films interface.
func (c client) Films() ([]Film, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, "movies/13/7925")

	r, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// Read body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	films := []Film{}

	err = json.Unmarshal(b, &films)
	if err != nil {
		return nil, err
	}

	return films, nil
}
