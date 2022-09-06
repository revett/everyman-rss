// Package rss provides HTTP handlers for RSS XML endpoints which Vercel will
// convert to serverless functions using the Go runtime.
// Note this cannot be within doc.go as Vercel sees that file as an endpoint.
package rss

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/revett/everyman-rss/internal/handler"
	"github.com/rs/zerolog/log"
)

// Films serves an RSS XML feed of the latest film releases from Everyman
// Cinema.
func Films(w http.ResponseWriter, r *http.Request) { //nolint:varnamelen
	log.Logger = commonLog.New()

	e := echo.New() //nolint:varnamelen
	e.Use(commonMiddleware.LoggerUsingZerolog(log.Logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			DisablePrintStack: true,
		},
	))

	e.GET("/*", handler.Films)
	e.ServeHTTP(w, r)
}
