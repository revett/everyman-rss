// Package handler (within directory called `api`) is a pattern that is enforced
// by Vercel for serverless functions using the Go runtime.
// See: https://vercel.com/docs/runtimes#official-runtimes/go
// Note this cannot be within doc.go as Vercel sees that file as an endpoint.
package handler

import (
	"context"
	_ "embed"
	"net/http"
	"text/template"

	"github.com/revett/everyman-rss/internal/api"
	"github.com/revett/everyman-rss/pkg/everyman"
)

//go:embed template/index.tmpl
var tmpl string

type templateCinemaValues struct {
	Name string
	Slug string
}

// Index serves a simple HTML page explaining the project.
func Index(w http.ResponseWriter, r *http.Request) {
	api.CommonMiddleware(index).ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := everyman.NewClientWithResponses(everyman.BaseWebURL)
	if err != nil {
		api.InternalServerError(
			w, err, "unable to create everyman api client",
		)
		return
	}

	cinemas, err := c.CinemasWithResponse(context.TODO())
	if err != nil {
		api.InternalServerError(
			w, err, "unable to request cinemas from everyman cinema api",
		)
		return
	}

	var templateData []templateCinemaValues
	for _, cinema := range *cinemas.JSON200 {
		templateData = append(
			templateData,
			templateCinemaValues{
				Name: cinema.CinemaName,
				Slug: cinema.Slug(),
			},
		)
	}

	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		api.InternalServerError(
			w, err, "failed to parse local template film",
		)
		return
	}

	err = t.Execute(w, templateData)
	if err != nil {
		api.InternalServerError(
			w, err, "failed when generating template for page",
		)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
