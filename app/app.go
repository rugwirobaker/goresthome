package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// Database driver
	_ "github.com/lib/pq"
	"github.com/rugwirobaker/structure/handlers"

	"github.com/gorilla/mux"
)

//App defines application wide configutration.s
type App struct {
	Router mux.Router
	DB     *sql.DB
}

//Initialize ...
func (a *App) Initialize() {
	fmt.Println("*** Initializing application...")
	a.initDb(Host, User, Password, Dbname, Port)
	a.initRoutes()
}

//Run method starts the server
func (a *App) Run(addr string) {
	fmt.Println("*** Starting the web server...")
	log.Fatal(http.ListenAndServe(addr, &a.Router))
}

//Routes definition

func (a *App) initRoutes() {
	a.Router.StrictSlash(false)

	a.Router.HandleFunc("/articles", func(w http.ResponseWriter,
		r *http.Request) {

		logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		fmt.Println(logger)

		handlers.RetrieveArticles(w, r, a.DB)
	}).Methods("GET")

	a.Router.HandleFunc("/articles/{id:[0-9]+}", func(w http.ResponseWriter,
		r *http.Request) {

		logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		fmt.Println(logger)

		handlers.RetrieveArticle(w, r, a.DB)
	}).Methods("GET")

	a.Router.HandleFunc("/articles", func(w http.ResponseWriter,
		r *http.Request) {

		logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		fmt.Println(logger)

		handlers.CreateArticle(w, r, a.DB)
	}).Methods("POST")
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
//TODO: find a way to move handlers to their own package
