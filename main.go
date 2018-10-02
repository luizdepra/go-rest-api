package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/pure"

	"github.com/luizdepra/go-rest-api/app"
	"github.com/luizdepra/go-rest-api/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := pure.New()

	app := app.New(router)
	app.RegisterRoutes()

	serverPort := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(serverPort, app.Router.Serve()))
}
