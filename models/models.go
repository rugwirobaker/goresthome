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

func retrieveArticle() error {
	return errors.New("Not implemented")
}

func createArticle(title, body, author string) error {
	return errors.New("Not implemented")
}

func listArticles() ([]Article, error) {
	return nil, errors.New("Not implemented")
}
