package models

import (
	"errors"
	"time"
)

//Article ...
type Article struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedOn time.Time `json:"createdon"`
}

//var id int

//Articles is represents an in memory database
var Articles = make(map[string]Article)

//RetrieveArticle ...
func (t *Article) RetrieveArticle() error {
	return errors.New("Not implemented")
}

//CreateArticle ...
func (t *Article) CreateArticle(title, body, author string) error {
	return errors.New("Not implemented")
}

//ListArticles ...
func ListArticles() []Article {
	var articles []Article
	for _, v := range Articles {
		articles = append(articles, v)
	}
	return articles
}
