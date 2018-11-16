package route

import (
	"net/http"

	"github.com/luizdepra/go-rest-api/app/model"
	"github.com/luizdepra/go-rest-api/app/repository"
)

// ListTasksHandler handles the Task listing request.
//
// GET /tasks/
func ListTasksHandler(repo *repository.TaskRepository, writer http.ResponseWriter, request *http.Request) {
	data, err := repo.List()
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	MakeJSONResponse(writer, http.StatusOK, data)
}

// CreateTaskHandler handles the Task creation request.
//
// POST /tasks/
func CreateTaskHandler(repo *repository.TaskRepository, writer http.ResponseWriter, request *http.Request) {
	var task model.Task

	err := decodeTask(request, &task)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid payload")
		return
	}

	data, err := repo.Create(&task)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	MakeJSONResponse(writer, http.StatusCreated, data)
}

// GetTaskHandler handles the Task retreave request.
//
// GET /tasks/:id
func GetTaskHandler(repo *repository.TaskRepository, writer http.ResponseWriter, request *http.Request, id int64) {
	data, err := repo.Get(id)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	if data == nil {
		MakeJSONErrorResponse(writer, http.StatusNotFound, "task not found")
		return
	}

	MakeJSONResponse(writer, http.StatusOK, data)
}

// UpdateTaskHandler handles the Task update request.
//
// PUT /tasks/:id
func UpdateTaskHandler(repo *repository.TaskRepository, writer http.ResponseWriter, request *http.Request, id int64) {
	var task model.Task

	err := decodeTask(request, &task)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusBadRequest, "invalid payload")
		return
	}

	task.ID = id

	data, err := repo.Update(&task)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	if data == nil {
		MakeJSONErrorResponse(writer, http.StatusNotFound, "task not found")
		return
	}

	MakeJSONResponse(writer, http.StatusOK, data)
}

// DeleteTaskHandler handles the Task deletion request.
//
// DELETE /tasks/:id
func DeleteTaskHandler(repo *repository.TaskRepository, writer http.ResponseWriter, request *http.Request, id int64) {
	data, err := repo.Delete(id)
	if err != nil {
		MakeJSONErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	if data == nil {
		MakeJSONErrorResponse(writer, http.StatusNotFound, "task not found")
		return
	}

	MakeJSONResponse(writer, http.StatusOK, data)
}
