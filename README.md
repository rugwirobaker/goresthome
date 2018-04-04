# Goresthome
In this repo I attempt to develop a general structure for REST APIs upon which
I will base all my future projects in go web development. Enjoy your cake.

## To run this application
* Set up a postgresql database with the details found in app/conf.go(or change them).
* Install the dependecies that we have listed here.
* Run "`goose up`" to migrate your database schema.
* Finally run "`go run main.go`" to start the server
* To test the API you can create some dummy data with [POSTMAN](https://www.getpostman.com/)

## Dependecies
* gorilla/mux: Install --> `go get -u github.com/gorilla/mux`
* lib/pq: Install --> `go get github.com/lib/pq`
* liamstask/goose: Install --> `go get bitbucket.org/liamstask/goose/cmd/goose`