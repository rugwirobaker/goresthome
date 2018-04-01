package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/structure/models"
)

// RetrieveArticle ...
func RetrieveArticle(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	title := urlParams["title"]
	//fmt.Println(title)
	article := models.RetrieveArticle(title)

	js, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// ListArticles ...
func ListArticles(w http.ResponseWriter, r *http.Request) {

	articles := models.ListArticles()
	//if err != nil {
	//panic(err)
	//}

	js, err := json.Marshal(articles)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

var id int

//CreateArticle ...
func CreateArticle(w http.ResponseWriter, r *http.Request) {

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
