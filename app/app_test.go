package app_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/pure"
	"github.com/stretchr/testify/assert"

	"github.com/luizdepra/go-rest-api/app"
)

func request(router *pure.Mux, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, body)
	recorder := httptest.NewRecorder()
	handler := router.Serve()
	handler.ServeHTTP(recorder, request)

	return recorder
}

func TestNew(t *testing.T) {
	router := pure.New()

	expected := &app.App{
		Router: router,
	}

	value := app.New(router)

	assert.Equal(t, expected, value)
}

func TestRegisterRoutes(t *testing.T) {
	router := pure.New()
	testApp := app.New(router)

	testApp.RegisterRoutes()

	response := request(testApp.Router, "GET", "/", nil)
	assert.Equal(t, 200, response.Code, "Root check failed")

	response = request(testApp.Router, "GET", "/tasks/", nil)
	assert.Equal(t, 200, response.Code, "Task List check failed")

	response = request(testApp.Router, "POST", "/tasks/", nil)
	assert.Equal(t, 200, response.Code, "Task Create check failed")

	response = request(testApp.Router, "GET", "/tasks/12345/", nil)
	assert.Equal(t, 200, response.Code, "Task Get check failed")

	response = request(testApp.Router, "PUT", "/tasks/12345/", nil)
	assert.Equal(t, 200, response.Code, "Task Update check failed")

	response = request(testApp.Router, "DELETE", "/tasks/12345/", nil)
	assert.Equal(t, 200, response.Code, "Task Delete check failed")

	response = request(testApp.Router, "GET", "/hello/", nil)
	assert.Equal(t, 404, response.Code, "Invalid route check failed")
}
