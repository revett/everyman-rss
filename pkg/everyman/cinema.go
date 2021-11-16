package everyman

import "github.com/gosimple/slug"

// Slug returns the cinema name as a slug.
func (c Cinema) Slug() string {
	return slug.Make(c.CinemaName)
}
