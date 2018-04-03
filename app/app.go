package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// Database driver
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/models"
)

//App the core struct of the aoo
type App struct {
	Router mux.Router
	DB     *sql.DB
}

//Initialize ...
func (a *App) Initialize() {
	a.initDb(Host, User, Password, Dbname, Port)
	a.initRoutes()
}

//Run method starts the server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, &a.Router))
}

//Routes definition

func (a *App) initRoutes() {
	//a.Router.HandleFunc("/articles", handlers.ListArticles).Methods("GET")
	//a.Router.HandleFunc("/articles/{title}", handlers.RetrieveArticle).Methods("GET")
	a.Router.HandleFunc("/articles", a.CreateArticle).Methods("POST")
}

//Initialize database connection
func (a *App) initDb(host, username, password, dbname string, port int) {
	//TODO: fix connection string --> DONE
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable", host, port, username, password, dbname)

	var err error

	a.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	//check if the datasource is valid
	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("*** Database Connection Established ...")
	}
}

//handlers

//CreateArticle ...
func (a *App) CreateArticle(w http.ResponseWriter, r *http.Request) {

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		panic(err)
	}

	article.CreateArticle(a.DB)

	js, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}
