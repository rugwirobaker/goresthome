package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/models"
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
	a.Router.HandleFunc("/articles", createArticle).Methods("POST")
}

//API represents dummy data
type API struct {
	//Message ...
	Message string `json:"message"`
}

//Last to be implemented
func getArticle(w http.ResponseWriter, r *http.Request) {
	message := API{"Hello,	world! I am an article"}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}

func listArticles(w http.ResponseWriter, r *http.Request) {

	articles := models.ListArticles()
	//if err != nil {
	//panic(err)
	//}

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(articles)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

var id int

func createArticle(w http.ResponseWriter, r *http.Request) {

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		panic(err)
	}

	models.CreateArticle(article)

	js, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}
