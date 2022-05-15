package main

import (
	"bytes"
	"context"
	_ "embed"
	"html/template"
	"os"

	commonLog "github.com/revett/common/log"
	"github.com/revett/everyman-rss/pkg/everyman"
	"github.com/rs/zerolog/log"
	"github.com/russross/blackfriday/v2"
)

var (
	//go:embed template/README.gen.md
	readmeMarkdown string

	//go:embed template/index.tmpl
	indexTemplate string
)

//go:generate cp ../../README.md template/README.gen.md
type templateData struct {
	README  template.HTML
	Cinemas []templateCinemaValues
}

type templateCinemaValues struct {
	Name string
	Slug string
}

func main() {
	log.Logger = commonLog.New()

	client, err := everyman.NewClientWithResponses(everyman.BaseWebURL)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create everyman api client")
	}

	cinemas, err := client.CinemasWithResponse(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg(
			"unable to request cinemas from everyman cinema api",
		)
	}

	markdown := blackfriday.Run(
		[]byte(readmeMarkdown),
	)

	templateData := templateData{
		README: template.HTML(markdown),
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
		log.Fatal().Err(err).Msg("failed to parse local template film")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		log.Fatal().Err(err).Msg("failed when generating page template")
	}

	file, err := os.Create("./build/index.html")
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create new file")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Info().Err(err).Msg("encountered error within defer statement")
		}
	}()

	_, err = file.WriteString(buf.String())
	if err != nil {
		log.Error().Err(err).Msg("failed to write template string to file")
	}
}
