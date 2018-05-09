//Package models Define interface for storage of different application components
package models

import (
	"database/sql"
	"errors"
	"time"
)

//Article models the database entity article
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedOn time.Time `json:"createdon"`
}

//ArticleResults models a list of retrieved articles
type ArticleResults struct {
	Count    int       `json:"article_count"`
	Articles []Article `json:"articles"`
}

//User models the database entity user
type User struct {
	ID         int       `json:"id,omitempty"`
	Fname      string    `json:"first_name"`
	Lname      string    `json:"last_name"`
	Email      string    `json:"email"`
	Pass       string    `json:"passwd"`
	PassHash   []byte    `json:"passwdhash,omitempty"`
	DateJoined time.Time `json:"joinedOn"`
}

//Users models a collection of user objects
type Users struct {
	Users []User `json:"users"`
	Count int    `json:"user_count"`
}

//var id int

//Articles is represents an in memory database
//var Articles = make(map[string]Article)

//RetrieveArticle ...
func (c *Article) RetrieveArticle(db *sql.DB) error {
	err := db.QueryRow("SELECT id, title, author FROM articles WHERE id=$1",
		c.ID).Scan(&c.ID, &c.Title, &c.Author)

	if err != nil {
		return err
	}
	return nil
}

//CreateArticle ...
func (c *Article) CreateArticle(db *sql.DB) error {
	//var article = Article{}
	c.CreatedOn = time.Now()
	err := db.QueryRow("INSERT INTO articles(title, body, author,createdon)"+
		"VALUES($1, $2, $3, $4)"+
		"RETURNING id", c.Title, c.Body, c.Author, c.CreatedOn).Scan(&c.ID)

	if err != nil {
		return err
	}
	return nil

}

//DeleteArticle query
func (c *Article) DeleteArticle(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM articles WHERE id=$1", c.ID)

	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleResults) getCount() {
	s.Count = len(s.Articles)
}

//ListArticles ...
func (s *ArticleResults) ListArticles(db *sql.DB) error {
	rows, err := db.Query("SELECT id, title, author FROM articles")

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var c Article
		if err := rows.Scan(&c.ID, &c.Title, &c.Author); err != nil {
			return err
		}
		s.Articles = append(s.Articles, c)
		s.getCount()
	}
	return nil
}

//RegisterUser creates an new user instance
func (u *User) RegisterUser(db *sql.DB) error {
	u.DateJoined = time.Now()
	err := db.QueryRow("INSERT INTO users(fname, lname, email,"+
		"passhash, datejoined) VALUES($1, $2, $3, $4, $5) RETURNING id",
		u.Fname, u.Lname, u.Email, u.PassHash, u.DateJoined).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser deletes user with dd:ID
func (u *User) DeleteUser(db *sql.DB) error {
	err := db.QueryRow("DELETE FROM users WHERE id=$1 "+
		"RETURNING id", u.ID).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

//UpdateUser update a user entity information
func (u *User) UpdateUser(db *sql.DB) error {
	return errors.New(" RetrieveUserByEmail Not implemented")
}

//RetrieveUserByEmail queries user with a given Email from the database
func (u *User) RetrieveUserByEmail(db *sql.DB) error {
	err := db.QueryRow("SELECT id, email, passhash FROM users WHERE email=$1",
		u.Email).Scan(&u.ID, &u.Email, &u.PassHash)

	if err != nil {
		return err
	}
	return nil
}

//ListUsers queries a list of user entities
func (us *Users) ListUsers(db *sql.DB) error {
	rows, err := db.Query("SELECT id, title, author FROM users")

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.DateJoined); err != nil {
			return err
		}
		us.Users = append(us.Users, u)
		us.getCount()
	}

	return nil
}

func (us *Users) getCount() {
	us.Count = len(us.Users)
}
