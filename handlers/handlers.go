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
type JSONResp struct {
	Status     string                 `json:"status"`
	Payload    *models.Article        `json:"payload,omitempty"`
	Results    *models.ArticleResults `json:"results,omitempty"`
	ErrMessage string                 `json:"error,omitempty"`
}

//CreateArticle ...
func CreateArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		panic(err)
	}

	article.CreateArticle(db)

	response := JSONResp{Status: "success", Payload: &article}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(js)
}

//RetrieveArticle ...
func RetrieveArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])

	if err != nil {
		log.Fatal(err)
	}

	article := models.Article{ID: id}
	article.RetrieveArticle(db)

	response := JSONResp{Status: "success", Payload: &article}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

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
	articles.ListArticles(db)

	response := JSONResp{Status: "success", Results: &articles}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
