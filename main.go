package main

import (
	"log"
	"net/http"

	"github.com/go-playground/pure"
)

func main() {
	p := pure.New()

	p.Get("/", rootHandler)

	todos := p.Group("/todos")
	todos.Get("", listTODOsHandler)
	todos.Post("/", createTodoHandler)
	todos.Get("/:id", getTODOHandler)
	todos.Put("/:id", updateTODOHandler)
	todos.Patch("/:id", updateTODOHandler)
	todos.Delete("/:id", deleteTODOHandler)

	log.Fatal(http.ListenAndServe(":8000", p.Serve()))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func listTODOsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listed!"))
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created!"))
}

func getTODOHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Got!"))
}

func updateTODOHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updated!"))
}

func deleteTODOHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleted!"))
}
