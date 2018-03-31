package models

import (
	"errors"
	"strconv"
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

var id int

//CreateArticle ...
func CreateArticle(article Article) Article {
	article.CreatedOn = time.Now()

	id++
	k := strconv.Itoa(id)
	Articles[k] = article
	return article
}

//ListArticles ...
func ListArticles() []Article {
	var articles []Article
	for _, v := range Articles {
		articles = append(articles, v)
	}
	return articles
}
