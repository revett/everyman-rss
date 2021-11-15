package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

// CommonMiddleware wraps a http.HandlerFunc in a set of common HTTP middleware
// handlers.
func CommonMiddleware(h http.HandlerFunc) http.Handler {
	loggingHandler := handlers.CombinedLoggingHandler(combinedLogger{}, h)
	return handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
		handlers.RecoveryLogger(recoveryHandlerLogger{}),
	)(loggingHandler)
}

type recoveryHandlerLogger struct{}

// Println implements the handlers.RecoveryHandlerLogger interface. It makes
// sure that log messages are consistently formatted.
func (r recoveryHandlerLogger) Println(v ...interface{}) {
	log.Printf("PNC %+v", v...)
}
