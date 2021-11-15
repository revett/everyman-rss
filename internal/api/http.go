package api

import (
	"net/http"
)

// BadRequest logs the error message, and writes the HTTP status code and error
// to the HTTP response.
func BadRequest(w http.ResponseWriter, err error, msg string) {
	logErrorAndWriteToResponse(w, err, msg, http.StatusBadRequest)
}

// InternalServerError logs the error message, and writes the HTTP status code
// and error to the HTTP response.
func InternalServerError(w http.ResponseWriter, err error, msg string) {
	logErrorAndWriteToResponse(w, err, msg, http.StatusInternalServerError)
}
