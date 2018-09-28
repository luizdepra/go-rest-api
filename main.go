package main

import (
	"log"
	"net/http"

	"github.com/go-playground/pure"

	"github.com/luizdepra/go-rest-api/app"
)

func main() {
	router := pure.New()

	app := app.New(router)
	app.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8000", app.Router.Serve()))
}
