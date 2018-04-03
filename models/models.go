package models

import (
	"database/sql"
	"log"
	"time"
)

//Article ...
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedOn time.Time `json:"createdon"`
}

//var id int

//Articles is represents an in memory database
var Articles = make(map[string]Article)

//RetrieveArticle ...
func RetrieveArticle(title string) Article {
	article := Article{}
	for _, a := range Articles {
		if a.Title == title {
			return a
		}
	}
	//fmt.Println("Could not retrieve Article")
	//TODO: Error logging
	return article
}

//CreateArticle ...
func (c *Article) CreateArticle(db *sql.DB) {
	//var article = Article{}
	c.CreatedOn = time.Now()
	err := db.QueryRow("INSERT INTO articles(title, body, author,createdon)"+
		"VALUES($1, $2, $3, $4)"+
		"RETURNING id", c.Title, c.Body, c.Author, c.CreatedOn).Scan(&c.ID)

	if err != nil {
		log.Fatal(err)
	}

}

//ListArticles ...
func ListArticles() []Article {
	var articles []Article
	for _, v := range Articles {
		articles = append(articles, v)
	}
	return articles
}
