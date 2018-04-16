package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
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

	n := negroni.New()
	n.Use(negroni.HandlerFunc(handlers.LoggingHandler))
	n.UseHandler(&a.Router)
	//log.Fatal(http.ListenAndServe(addr, &a.Router))
	n.Run(addr)
}

//Routes definition

func (a *App) initRoutes() {
	a.Router.StrictSlash(false)

	// Route: Retrieve an article
	a.Router.HandleFunc("/articles/{id:[0-9]+}", func(w http.ResponseWriter,
		r *http.Request) {

		//logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		//fmt.Println(logger)

		handlers.RetrieveArticle(w, r, a.DB)
	}).Methods("GET")

	// Route: Creates an article
	a.Router.HandleFunc("/articles", func(w http.ResponseWriter,
		r *http.Request) {

		//logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		//fmt.Println(logger)

		handlers.CreateArticle(w, r, a.DB)
	}).Methods("POST")

	// Route: Deletes an article
	a.Router.HandleFunc("/articles/{id:[0-9]+}", func(w http.ResponseWriter,
		r *http.Request) {

		//logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		//fmt.Println(logger)

		handlers.DeleteArticle(w, r, a.DB)
	}).Methods("DELETE")

	// Route: Retrieves a list of articles
	a.Router.HandleFunc("/articles", func(w http.ResponseWriter,
		r *http.Request) {

		//logger := fmt.Sprintf("*** request: %s | %s%s", r.Method, r.Host, r.URL)
		//fmt.Println(logger)

		handlers.RetrieveArticles(w, r, a.DB)
	}).Methods("GET")

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
