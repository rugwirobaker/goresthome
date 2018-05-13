//Package routes defines web application routes.
package routes

import (
	"crypto/rsa"
	"database/sql"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/handlers"
)

//InitArticleRoutes initializes user resource routes
func InitArticleRoutes(router *mux.Router, db *sql.DB, vKey *rsa.PublicKey) *mux.Router {
	artRoutes := mux.NewRouter().StrictSlash(false)
	artRoutes.HandleFunc("/articles/{id:[0-9]+}", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RetrieveArticle(w, r, db)
	}).Methods("GET")

	// Route: Creates an article
	artRoutes.HandleFunc("/articles", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.CreateArticle(w, r, db)
	}).Methods("POST")

	// Route: Deletes an article
	artRoutes.HandleFunc("/articles/{id:[0-9]+}", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.DeleteArticle(w, r, db)
	}).Methods("DELETE")

	// Route: Retrieves a list of articles
	artRoutes.HandleFunc("/articles", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RetrieveArticles(w, r, db)
	}).Methods("GET")

	router.PathPrefix("/articles").Handler(negroni.New(
		negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			handlers.AuthHandler(w, r, next, vKey)
		}),
		negroni.Wrap(artRoutes),
	))
	return router
}

//InitUserRoutes initializes user resource routes
func InitUserRoutes(router *mux.Router, db *sql.DB, sKey *rsa.PrivateKey) *mux.Router {
	userRoutes := mux.NewRouter().StrictSlash(false)
	//Handlers new user registration
	userRoutes.HandleFunc("/users/register", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RegisterUser(w, r, db)
	}).Methods("POST")

	//Handlers user login
	userRoutes.HandleFunc("/users/login", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.LoginUser(w, r, db, sKey)
	}).Methods("POST")

	//Handlers user delete
	userRoutes.HandleFunc("/users/delete", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RegisterUser(w, r, db)
	}).Methods("DELETE")

	//Handlers user list i.e returns user collection
	userRoutes.HandleFunc("/users", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RetrieveUsers(w, r, db)
	}).Methods("GET")

	router.PathPrefix("/users").Handler(negroni.New(
		negroni.Wrap(userRoutes),
	))
	return router
}
