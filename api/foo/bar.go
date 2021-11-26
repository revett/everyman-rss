package foo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/rs/zerolog/log"
)

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
	return ctx.String(http.StatusOK, "test")
}
