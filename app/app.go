package app

import (
	"github.com/go-playground/pure"

	"github.com/luizdepra/go-rest-api/app/route"
)

// App holds common API structures.
type App struct {
	Router *pure.Mux
}

// New creates and returns a new App instance.
func New(router *pure.Mux) *App {
	return &App{
		Router: router,
	}
}

// RegisterRoutes registers all availlable routes into the router.
func (a *App) RegisterRoutes() {
	a.Router.Get("/", route.RootHandler)

	tasks := a.Router.Group("/tasks")
	tasks.Get("/", route.ListTasksHandler)
	tasks.Post("/", route.CreateTaskHandler)
	tasks.Get("/:id/", route.GetTaskHandler)
	tasks.Put("/:id/", route.UpdateTaskHandler)
	tasks.Delete("/:id/", route.DeleteTaskHandler)
}
