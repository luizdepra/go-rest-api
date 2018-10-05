package route

import (
	"net/http"

	"github.com/luizdepra/go-rest-api/app/model"
)

// ListTasksHandler handles the Task listing request.
//
// GET /tasks/
func ListTasksHandler(writer http.ResponseWriter, request *http.Request) {
	payload := getAllTasks()

	MakeJSONResponse(writer, http.StatusOK, payload)
}

// CreateTaskHandler handles the Task creation request.
//
// POST /tasks/
func CreateTaskHandler(writer http.ResponseWriter, request *http.Request) {
	var task model.Task

	err := decodeTask(request, &task)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid payload")
		return
	}

	addTask(&task)

	MakeJSONResponse(writer, http.StatusCreated, task)
}

// GetTaskHandler handles the Task retreave request.
//
// GET /tasks/:id
func GetTaskHandler(writer http.ResponseWriter, request *http.Request, id int64) {
	task := getTask(id)
	if task == nil {
		MakeJSONErrorResponse(writer, http.StatusNotFound, "task not found")
		return
	}

	MakeJSONResponse(writer, http.StatusOK, task)
}

// UpdateTaskHandler handles the Task update request.
//
// PUT /tasks/:id
func UpdateTaskHandler(writer http.ResponseWriter, request *http.Request, id int64) {
	var task model.Task

	err := decodeTask(request, &task)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid payload")
		return
	}

	updatedTask := updateTask(id, &task)
	if updatedTask == nil {
		MakeJSONErrorResponse(writer, http.StatusNotFound, "task not found")
		return
	}

	MakeJSONResponse(writer, http.StatusOK, updatedTask)
}

// DeleteTaskHandler handles the Task deletion request.
//
// DELETE /tasks/:id
func DeleteTaskHandler(writer http.ResponseWriter, request *http.Request, id int64) {
	task := deleteTask(id)
	if task == nil {
		MakeJSONErrorResponse(writer, http.StatusNotFound, "task not found")
		return
	}

	MakeJSONResponse(writer, http.StatusOK, task)
}
