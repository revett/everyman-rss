package everyman

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

// Film is the struct representation of a film.
type Film struct {
	ID           int            `json:"FilmId"`
	Title        string         `json:"Title"`
	Certificate  string         `json:"Cert"`
	Img          string         `json:"Img"`
	Trailer      string         `json:"Trailer"`
	ReleaseDate  string         `json:"ReleaseDate"`
	RunTime      string         `json:"RunTime"`
	Synopsis     string         `json:"Synopsis"`
	Teaser       string         `json:"Teaser"`
	Cast         string         `json:"Cast"`
	Director     string         `json:"Director"`
	FriendlyName string         `json:"FriendlyName"`
	Order        int            `json:"Order"`
	MediaItems   FilmMediaItems `json:"MediaItems"`
}

// FilmMediaItems is the struct representation of media links that relate to a
// film.
type FilmMediaItems struct {
	YouTubeTrailer string `json:"YouTubeTrailer"`
	QuadStill      string `json:"QuadStill"`
	Trailer        string `json:"Trailer"`
	OneSheet       string `json:"OneSheet"`
}

// Description generates a text description for the film, with an image URL
// appended to the end if one is present. The description is in HTML.
func (f Film) Description() string {
	d := f.Synopsis

	if !f.HasImage() {
		return d
	}

	return fmt.Sprintf(
		`%s - <img src="%s" />`, d, f.Image(),
	)
}

// HasImage checks if the film has an image.
func (f Film) HasImage() bool {
	return f.Img != "" || f.MediaItems.QuadStill != ""
}

// IDStr returns the Film ID as a string instead of an int.
func (f Film) IDStr() string {
	return strconv.Itoa(f.ID)
}

// Image returns the best available image for the film.
func (f Film) Image() string {
	if f.MediaItems.QuadStill != "" {
		return f.MediaItems.QuadStill
	}

	if f.Img != "" {
		return f.Img
	}

	return ""
}

// ImageLength returns the size of the image in bytes.
func (f Film) ImageLength() (int64, error) {
	if !f.HasImage() {
		return 0, nil
	}

	r, err := http.Head(
		f.Image(),
	)
	if err != nil {
		return 0, err
	}

	cl := r.Header.Get("Content-Length")
	if cl == "" {
		return 0, nil
	}

	size, err := strconv.Atoi(cl)
	if err != nil {
		return 0, err
	}

	return int64(size), nil
}

// ImageMIMEType returns the MIME type for the image of the film.
func (f Film) ImageMIMEType() string {
	ext := strings.ToLower(
		filepath.Ext(
			f.Image(),
		),
	)

	switch ext {
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	}

	return ""
}

// URL generates the webpage URL for the film.
func (f Film) URL() string {
	return fmt.Sprintf("%s/film-info/%s", BaseWebURL, f.FriendlyName)
}
