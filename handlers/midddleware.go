package handlers

import (
	"fmt"
	"log"
	"net/http"
)

//LoggingHandler ...
func LoggingHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
	log.Println(logger)

	//set application wide http response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	next(w, r)
}

func cachingHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}
