package route

import "net/http"

// ListTasksHandler handles the Task listing request.
//
// GET /tasks/
func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listed!"))
}

// CreateTaskHandler handles the Task creation request.
//
// POST /tasks/
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created!"))
}

// GetTaskHandler handles the Task retreave request.
//
// GET /tasks/:id
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Got!"))
}

// UpdateTaskHandler handles the Task update request.
//
// PUT /tasks/:id
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updated!"))
}

// DeleteTaskHandler handles the Task deletion request.
//
// DELETE /tasks/:id
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleted!"))
}
