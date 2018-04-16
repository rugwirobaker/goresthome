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

	next(w, r)
}

func cachingHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}
