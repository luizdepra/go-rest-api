package route

import (
	"encoding/json"
	"net/http"

	"github.com/luizdepra/go-rest-api/app/model"
)

// MakeJSONResponse creates a response with a payload in JSON format.
func MakeJSONResponse(writer http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte(response))
}

// MakeJSONErrorResponse creates a error response with a message in JSON format.
func MakeJSONErrorResponse(writer http.ResponseWriter, status int, message string) {
	MakeJSONResponse(writer, status, map[string]string{"message": message})
}

// decodeTask creates a Task from  a request.Body.
func decodeTask(request *http.Request, task *model.Task) error {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&task); err != nil {
		return err
	}

	return nil
}

// getAllTasks gets all Tasks.
func getAllTasks() []*model.Task {
	return model.TaskRepository
}

// getTask gets a Task with the given id.
func getTask(id int64) *model.Task {
	for _, v := range model.TaskRepository {
		if v.ID == id {
			return v
		}
	}

	return nil
}

// addTask adds a new Task.
func addTask(task *model.Task) {
	model.TaskRepository = append(model.TaskRepository, task)
}

// updateTask updates a Task with new values.
func updateTask(id int64, task *model.Task) *model.Task {
	original := getTask(id)
	if original == nil {
		return nil
	}

	original.Title = task.Title
	original.Priority = task.Priority
	original.Done = task.Done

	return original
}

// deleteTask removes a Task with the given id.
func deleteTask(id int64) *model.Task {
	task := getTask(id)
	if task == nil {
		return nil
	}

	list := []*model.Task{}
	for _, v := range model.TaskRepository {
		if v.ID != id {
			list = append(list, v)
		}
	}
	model.TaskRepository = list

	return task
}
