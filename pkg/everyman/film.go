package everyman

import "fmt"

// URL returns the webpage URL for the film.
func (f Film) URL() string {
	return fmt.Sprintf("%s/film-info/%s", BaseWebURL, f.FriendlyName)
}
