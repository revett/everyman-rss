package everyman

import (
	"fmt"
	"strconv"
)

// Film is the struct representation of a film.
type Film struct {
	ID           int            `json:"FilmId"`
	Title        string         `json:"Title"`
	Certificate  string         `json:"Cert"`
	Image        string         `json:"Img"`
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

const baseWebURL = "https://www.everymancinema.com"

// IDStr returns the Film ID as a string instead of an int.
func (f Film) IDStr() string {
	return strconv.Itoa(f.ID)
}

func (f Film) URL() string {
	return fmt.Sprintf("%s/film-info/%s", baseWebURL, f.FriendlyName)
}
