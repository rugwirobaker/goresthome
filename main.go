package main

import "github.com/rugwirobaker/structure/app"

func main() {
	a := &app.App{}
	a.Initialize()
	a.Run(":8080")
}

//TODO: move handlers to own package
//TODO: improve error reporting
//TODO: add caching middleware
//TODO: add users and authentication
