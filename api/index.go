// Package handler (within directory called `api`) is a pattern that is enforced
// by Vercel for serverless functions using the Go runtime.
// See: https://vercel.com/docs/runtimes#official-runtimes/go
// Note this cannot be within doc.go as Vercel sees that file as an endpoint.
package handler

import (
	_ "embed"
	"net/http"
	"text/template"

	"github.com/revett/everyman-rss/internal/api"
)

//go:embed template/index.tmpl
var foo string

// Index serves a plaintext string with a link to the Github repo.
func Index(w http.ResponseWriter, r *http.Request) {
	api.CommonMiddleware(index).ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").Parse(foo)
	if err != nil {
		api.InternalServerError(
			w, err, "failed to parse local template film",
		)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	err = t.Execute(w, nil)
	if err != nil {
		api.InternalServerError(
			w, err, "failed when generating template for page",
		)
		return
	}
}
