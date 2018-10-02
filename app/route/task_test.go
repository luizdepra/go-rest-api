package route_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/luizdepra/go-rest-api/app/route"
)

func TestListTasksHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/tasks/", nil)

	route.ListTasksHandler(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Listed!", recorder.Body.String())
}

func TestCreateTaskHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/tasks/", nil)

	route.CreateTaskHandler(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Created!", recorder.Body.String())
}

func TestGetTaskHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/tasks/12345/", nil)

	route.GetTaskHandler(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Got!", recorder.Body.String())
}

func TestUpdateTaskHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("PUT", "/tasks/12345/", nil)

	route.UpdateTaskHandler(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Updated!", recorder.Body.String())
}

func TestDeleteTaskHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("DELETE", "/tasks/12345/", nil)

	route.DeleteTaskHandler(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Deleted!", recorder.Body.String())
}
