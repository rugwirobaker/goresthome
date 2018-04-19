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

	err = article.CreateArticle(db)
	if err != nil {

		Err := StatusError{Err: err, Code: http.StatusBadRequest,
			Message: "bad request"}
		// Test for HttpStatus in loggingMiddleware
		//log the error(StatusMessage)

		response := JSONResp{Status: "fail", ErrMessage: Err.Message}
		js, _ := json.Marshal(response)
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(Err.Code)
		w.Write(js)
	} else {
		response := JSONResp{Status: "success", Payload: &article}
		js, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		w.Write(js)
	}
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
		Err := StatusError{Err: err, Code: http.StatusNotFound,
			Message: "resource not found"}
		// Test for HttpStatus in loggingMiddleware
		//log the error(StatusMessage)

		response := JSONResp{Status: "fail", ErrMessage: Err.Message}
		js, _ := json.Marshal(response)
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(Err.Code)
		w.Write(js)
	} else {
		response := JSONResp{Status: "success", Payload: &article}
		js, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	}

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
		response := JSONResp{Status: "fail", ErrMessage: err.Error()}
		js, _ := json.Marshal(response)
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(js)
	} else {

		response := JSONResp{Status: "success", Results: &articles}
		js, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

//DeleteArticle ...
func DeleteArticle(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])

	if err != nil {
		log.Fatal(err)
	}
	article := models.Article{ID: id}
	err = article.DeleteArticle(db)

	if err != nil {
		response := JSONResp{Status: "fail", ErrMessage: err.Error()}
		js, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(js)
	} else {
		response := JSONResp{Status: "success"}
		js, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(js)

	}
}
