package app

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-playground/pure"

	"github.com/luizdepra/go-rest-api/app/repository"
	"github.com/luizdepra/go-rest-api/app/route"
)

// App holds common API structures.
type App struct {
	Router     *pure.Mux
	Repository *repository.TaskRepository
}

// New creates and returns a new App instance.
func New(router *pure.Mux, db *sql.DB) *App {
	repo := repository.New(db)

	return &App{
		Router:     router,
		Repository: repo,
	}
}

// RegisterRoutes registers all availlable routes into the router.
func (app *App) RegisterRoutes() {
	app.Router.Get("/", app.GetRoot)

	tasks := app.Router.Group("/tasks")
	tasks.Get("/", app.ListTasks)
	tasks.Post("/", app.CreateTask)
	tasks.Get("/:id/", app.GetTask)
	tasks.Put("/:id/", app.UpdateTask)
	tasks.Delete("/:id/", app.DeleteTask)
}

// parseID parses the Id from the Request.
func (app *App) parseID(request *http.Request) (int64, error) {
	requestVars := pure.RequestVars(request)
	idParam := requestVars.URLParam("id")
	return strconv.ParseInt(idParam, 10, 64)
}

// GetRoot warps the RootHandler.
func (app *App) GetRoot(writer http.ResponseWriter, request *http.Request) {
	route.RootHandler(writer, request)
}

// ListTasks wraps the ListTasksHandler.
func (app *App) ListTasks(writer http.ResponseWriter, request *http.Request) {
	route.ListTasksHandler(app.Repository, writer, request)
}

// CreateTask wraps the CreateTaskHandler.
func (app *App) CreateTask(writer http.ResponseWriter, request *http.Request) {
	route.CreateTaskHandler(app.Repository, writer, request)
}

// GetTask wraps the GetTaskHandler.
func (app *App) GetTask(writer http.ResponseWriter, request *http.Request) {
	id, err := app.parseID(request)
	if err != nil {
		route.MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid id value")
		return
	}

	route.GetTaskHandler(app.Repository, writer, request, id)
}

// UpdateTask wraps the UpdateTaskHandler.
func (app *App) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	id, err := app.parseID(request)
	if err != nil {
		route.MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid id value")
		return
	}

	route.UpdateTaskHandler(app.Repository, writer, request, id)
}

// DeleteTask wraps the DeleteTaskHandler.
func (app *App) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	id, err := app.parseID(request)
	if err != nil {
		route.MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid id value")
		return
	}

	route.DeleteTaskHandler(app.Repository, writer, request, id)
}
