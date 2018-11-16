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
