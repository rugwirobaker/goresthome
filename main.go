package main

import "github.com/rugwirobaker/structure/app"

func main() {
	a := app.App{}
	a.Initialize()
	a.Run(":8080")
}
