// Package handler (within directory called `api`) is a pattern that is enforced
// by Vercel for serverless functions using the Go runtime.
// See: https://vercel.com/docs/runtimes#official-runtimes/go
// Note this cannot be within doc.go as Vercel sees that file as an endpoint.
package handler

import (
	"fmt"
	"net/http"

	"github.com/revett/everyman-rss/internal/api"
)

// Index serves a plaintext string with a link to the Github repo.
func Index(w http.ResponseWriter, r *http.Request) {
	api.CommonMiddleware(index).ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(
		w, "Read about the project: https://github.com/revett/everyman-rss",
	)
}
