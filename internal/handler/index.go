package handler

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/russross/blackfriday/v2"
)

// Index serves a simple HTML page explaining the project.
func Index(ctx echo.Context) error {
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

	if err := ctx.HTML(http.StatusOK, buf.String()); err != nil {
		return fmt.Errorf("failed to send html with status code: %w", err)
	}

	return nil
}

var (
	//go:embed template/README.gen.md
	readmeMarkdown string

	//go:embed template/index.tmpl
	indexTemplate string
)

//go:generate cp ../../README.md template/README.gen.md
type templateData struct {
	README  string
	Cinemas []templateCinemaValues
}

type templateCinemaValues struct {
	Name string
	Slug string
}
