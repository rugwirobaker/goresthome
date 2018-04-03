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
//var Articles = make(map[string]Article)

//RetrieveArticle ...
func (c *Article) RetrieveArticle(db *sql.DB) {
	err := db.QueryRow("SELECT id, title, author FROM articles WHERE id=$1",
		c.ID).Scan(&c.ID, &c.Title, &c.Author)

	if err != nil {
		log.Fatal(err)
	}
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
func ListArticles(db *sql.DB) []Article {
	var articles []Article
	rows, err := db.Query("SELECT id, title, author FROM articles")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Article
		if err := rows.Scan(&c.ID, &c.Title, &c.Author); err != nil {
			log.Fatal(err)
		}
		articles = append(articles, c)
	}
	return articles
}
