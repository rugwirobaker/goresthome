package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/models"
)

//JSONResp is the response structure
//type JSONResp struct {
//	Status     string                 `json:"status"`
//	Payload    *models.Article        `json:"payload,omitempty"`
//	Results    *models.ArticleResults `json:"results,omitempty"`
//	ErrMessage string                 `json:"error,omitempty"`
//}

//CreateArticle ...
func CreateArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	err = article.CreateArticle(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, article)
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
		respondWithError(w, http.StatusBadRequest, "Invalid object ID")
	}

	respondWithJSON(w, http.StatusOK, article)
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
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, articles)
}

//DeleteArticle ...
func DeleteArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid object ID")
	}
	article := models.Article{ID: id}
	err = article.DeleteArticle(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//helper functions

//respondWithJSON responds with json encoded data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

//respondWithError respond with an error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
