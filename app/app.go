package app

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mediocregopher/radix.v2/pool"

	"github.com/codegangsta/negroni"
	// Database driver
	_ "github.com/lib/pq"
	"github.com/rugwirobaker/structure/handlers"

	"github.com/gorilla/mux"
)

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "./keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "./keys/app.rsa.pub"
)

//Private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

//App defines application wide configutration.s
type App struct {
	Router    mux.Router
	DB        *sql.DB
	Store     *pool.Pool
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
}

//Initialize ...
func (a *App) Initialize() {
	fmt.Println("*** Initializing application...")
	a.initKeys()
	a.initDb(Host, User, Password, Dbname, Port)
	a.initRoutes()
}

//Run method starts the server
func (a *App) Run(addr string) {
	fmt.Println("*** Starting the web server...")

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(handlers.LoggingHandler),
	)
	//n.Use(negroni.HandlerFunc(handlers.LoggingHandler))
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

	///////////////////////////////////////////////////////////////////////////////
	//userHandlers

	//Handlers new user registration
	a.Router.HandleFunc("/users/register", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RegisterUser(w, r, a.DB)
	}).Methods("POST")

	//Handlers user login
	a.Router.HandleFunc("/users/login", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.LoginUser(w, r, a.DB, a.signKey)
	}).Methods("POST")

	//Handlers user delete
	a.Router.HandleFunc("/users/delete", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RegisterUser(w, r, a.DB)
	}).Methods("DELETE")

	//Handlers user list i.e returns user collection
	a.Router.HandleFunc("/users/", func(w http.ResponseWriter,
		r *http.Request) {
		handlers.RetrieveUsers(w, r, a.DB)
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

//initStore initializes the redis database
func (a *App) initStore(host, port string) {}

//initKeys loads the application rsa keys
func (a *App) initKeys() {
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	a.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	a.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}
