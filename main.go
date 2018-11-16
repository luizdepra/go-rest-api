package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/pure"
	_ "github.com/mattn/go-sqlite3"

	"github.com/luizdepra/go-rest-api/app"
	"github.com/luizdepra/go-rest-api/config"
	"github.com/luizdepra/go-rest-api/middleware"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := pure.New()
	router.Use(middleware.Logging())

	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := app.New(router, db)
	app.RegisterRoutes()

	serverPort := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(serverPort, app.Router.Serve()))
}
