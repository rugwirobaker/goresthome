//Package app defines web application-wide initial configuration.
package app

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/codegangsta/negroni"
	// Database driver
	_ "github.com/lib/pq"
	"github.com/rugwirobaker/structure/handlers"
	"github.com/rugwirobaker/structure/routes"

	"github.com/gorilla/mux"
)

//Private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

//App defines application wide configutration.s
type App struct {
	Conf      *Config
	Router    *mux.Router
	DB        *sql.DB
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
}

//Initialize ...
func (a *App) Initialize() {
	fmt.Println("*** Initializing application...")
	a.initKeys()
	a.initDb(a.Conf.DB.Host,
		a.Conf.DB.Username,
		a.Conf.DB.Password,
		a.Conf.DB.Dbname,
		a.Conf.DB.Port)
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
	n.UseHandler(a.Router)
	//log.Fatal(http.ListenAndServe(addr, &a.Router))
	n.Run(addr)
}

//Routes initializtion
func (a *App) initRoutes() {
	a.Router = mux.NewRouter().StrictSlash(false)
	a.Router = routes.InitArticleRoutes(a.Router, a.DB, a.verifyKey)
	a.Router = routes.InitUserRoutes(a.Router, a.DB, a.signKey)
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

	signBytes, err := ioutil.ReadFile(a.Conf.Keys.SignKey)
	fatal(err)

	a.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(a.Conf.Keys.VerifyKey)
	fatal(err)

	a.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}
