//Package handlers defines route handlers.
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rugwirobaker/structure/security"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/models"
)

//messages
var (
	fail    = "failure"
	success = "success"
)

type (
	loginData struct {
		ID     int    `json:"id,omitempty"`
		Email  string `json:"email"`
		Passwd string `json:"password"`
		Token  string `json:"token"`
	}

	registrationData struct {
		ID     int       `json:"id,omitempty"`
		Fname  string    `json:"first_name"`
		Lname  string    `json:"last_name"`
		Email  string    `json:"email"`
		Passwd string    `json:"passwd"`
		Joined time.Time `json:"joined"`
	}
)

//CreateArticle ...
func CreateArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fail, "Invalid request payload")
	}

	err = article.CreateArticle(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fail, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, success, article)
}

//RetrieveArticle ...
func RetrieveArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])

	if err != nil {
		log.Fatal(err)
	}

	article := models.Article{ID: id}
	err = article.RetrieveArticle(db)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fail, "Invalid object ID")
	}

	respondWithJSON(w, http.StatusOK, success, article)
}

//RetrieveArticles ...
func RetrieveArticles(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//count, _ := strconv.Atoi(r.FormValue("count"))
	//start, _ := strconv.Atoi(r.FormValue("start"))

	//if count > 10 || count < 1 {
	//count = 10
	//}
	//if start < 0 {
	//start = 0
	//}
	var articles models.ArticleResults
	err := articles.ListArticles(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fail, err.Error())
	}

	respondWithJSON(w, http.StatusOK, success, articles)
}

//DeleteArticle ...
func DeleteArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fail, "Invalid object ID")
	}
	article := models.Article{ID: id}
	err = article.DeleteArticle(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fail, err.Error())
	}

	respondWithJSON(w, http.StatusOK, success, map[string]string{"result": success})
}

//RegisterUser endpoint creates a new user account
//The required user data is defined in models/User
func RegisterUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//respondWithError(w, http.StatusInternalServerError, "Not implemented")
	var user registrationData

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fail, "Invalid request payload")
	}

	var usermodel models.User

	hash, err := security.HashPassword([]byte(user.Passwd))
	if err != nil {
		log.Fatalf(err.Error())
	}

	usermodel = models.User{
		Fname:    user.Fname,
		Lname:    user.Lname,
		Email:    user.Email,
		Pass:     user.Passwd,
		PassHash: hash,
	}

	err = usermodel.RegisterUser(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fail, err.Error())
	} else {
		user.ID = usermodel.ID
		user.Joined = usermodel.DateJoined
		respondWithJSON(w, http.StatusCreated, success, user)
	}
}

//LoginUser endpoint requires a email and password for login
func LoginUser(w http.ResponseWriter, r *http.Request, db *sql.DB, signKey interface{}) {
	//respondWithError(w, http.StatusInternalServerError, "Not implemented")
	var user loginData

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fail, "Invalid request payload")
	}

	usermodel := models.User{Email: user.Email}

	err = usermodel.RetrieveUserByEmail(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fail, err.Error())
	} else {
		if security.CheckPasswordHash(usermodel.PassHash, user.Passwd) {
			var token string
			token, err = security.GenerateJWT(usermodel.Email, "admin", signKey)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, fail, err.Error())
			} else {
				user.Token = token
				user.ID = usermodel.ID
				respondWithJSON(w, http.StatusOK, success, user)
			}
		} else {
			respondWithError(w, http.StatusNotFound, fail, "invalid credentials")
		}
	}
}

//DeleteUser deletes a user with a given email and requires authentication
func DeleteUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	respondWithError(w, http.StatusInternalServerError, fail, "Not implemented")
}

//RetrieveUser retrieves a given user profile
func RetrieveUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	respondWithError(w, http.StatusInternalServerError, fail, "Not implemented")
}

//RetrieveUsers retrievies a list of users
func RetrieveUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	respondWithError(w, http.StatusInternalServerError, fail, "Not implemented")
}
