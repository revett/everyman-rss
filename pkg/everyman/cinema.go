package everyman

import (
	"fmt"

	"github.com/gosimple/slug"
)

// Slug returns the cinema name as a slug.
func (c Cinema) Slug() string {
	return slug.Make(c.CinemaName)
}

// URL returns the web URL to the cinema page.
func (c Cinema) URL() string {
	return fmt.Sprintf("%s/%s", BaseWebURL, c.Slug())
}
