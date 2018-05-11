package handlers

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//LogResponseWriter wraps the builtin net/http ResponseWriter
// for response logging purposes.
type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

//NewLogResponseWriter creates a new loggingResponse type.
func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{w, http.StatusOK}
}

//WriteHeader initializes loggingResponseWriter values.
func (lrw *LogResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

//LoggingHandler ...
func LoggingHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
	log.Println(logger)

	lrw := NewLogResponseWriter(w)

	//set application wide http response headers
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

	next(lrw, r)

	statusCode := lrw.statusCode
	logger = fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode))
	log.Println(logger)

}

//AuthHandler validates a request's jwt token to access a protected resource.
func AuthHandler(w http.ResponseWriter, r *http.Request,
	next http.HandlerFunc, vKey *rsa.PublicKey) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return vKey, nil
		})
	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			respondWithError(w, http.StatusUnauthorized, "Token is not valid")
		}
	} else {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized access to this resource")
	}
}

//CachingHandler uses redis/memcache db to cache content
func CachingHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}

//func setHeadersHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}
