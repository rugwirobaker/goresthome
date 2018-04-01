package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/handlers"
)

//App the core struct of the aoo
type App struct {
	Router mux.Router
	DB     sql.DB
}

//Initialize ...
func (a *App) Initialize() {
	a.initRoutes()
}

//Run method starts the server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, &a.Router))
}

//Routes definition

func (a *App) initRoutes() {
	a.Router.HandleFunc("/articles", handlers.ListArticles).Methods("GET")
	a.Router.HandleFunc("/articles/{title}", handlers.RetrieveArticle).Methods("GET")
	a.Router.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")
}
