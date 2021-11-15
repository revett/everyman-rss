package api

import (
	"log"
	"net/http"
)

func logErrorAndWriteToResponse(w http.ResponseWriter, err error, msg string, code int) {
	log.Printf("ERR %s", err.Error())
	http.Error(w, msg, code)
}
