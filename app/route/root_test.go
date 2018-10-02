package route_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/luizdepra/go-rest-api/app/route"
)

func TestRootHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	route.RootHandler(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Root root!", recorder.Body.String())
}
