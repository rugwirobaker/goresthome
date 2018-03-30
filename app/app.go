package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	a.Router.HandleFunc("/articles", listArticles).Methods("GET")
	a.Router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	a.Router.HandleFunc("/articles/", createArticle).Methods("POST")
}

//API represents dummy data
type API struct {
	//Message ...
	Message string `json:"message"`
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	message := API{"Hello,	world! I am an article"}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	message := API{"Hello,	world! This is a list"}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	message := API{"Hello,	world! I create articles"}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}
