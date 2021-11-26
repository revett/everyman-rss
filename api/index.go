// Package handler (within directory called `api`) is a pattern that is enforced
// by Vercel for serverless functions using the Go runtime.
// See: https://vercel.com/docs/runtimes#official-runtimes/go
// Note this cannot be within doc.go as Vercel sees that file as an endpoint.
package api

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/rs/zerolog/log"
	"github.com/russross/blackfriday/v2"
)

var (
	//go:embed template/readme.tmpl.md
	readmeMarkdown string

	//go:embed template/index.tmpl
	indexTemplate string
)

//go:generate cp ../README.md template/readme.tmpl.md
type templateData struct {
	README  string
	Cinemas []templateCinemaValues
}

type templateCinemaValues struct {
	Name string
	Slug string
}

// Index serves a simple HTML page explaining the project.
func Index(w http.ResponseWriter, r *http.Request) { // nolint:varnamelen
	log.Logger = commonLog.New()

	e := echo.New() // nolint:varnamelen
	e.Use(commonMiddleware.LoggerUsingZerolog(log.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.GET("/*", indexHandler)
	e.ServeHTTP(w, r)
}

func indexHandler(ctx echo.Context) error {
	client, err := everyman.NewClientWithResponses(everyman.BaseWebURL)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "unable to create everyman api client",
		)
	}

	cinemas, err := client.CinemasWithResponse(context.TODO())
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"unable to request cinemas from everyman cinema api",
		)
	}

	templateData := templateData{
		README: string(
			blackfriday.Run(
				[]byte(readmeMarkdown),
			),
		),
	}

	for _, cinema := range *cinemas.JSON200 {
		templateData.Cinemas = append(
			templateData.Cinemas,
			templateCinemaValues{
				Name: cinema.CinemaName,
				Slug: cinema.Slug(),
			},
		)
	}

	tmpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "failed to parse local template film",
		)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, "failed when generating page template",
		)
	}

	return ctx.HTML(http.StatusOK, buf.String())
}
