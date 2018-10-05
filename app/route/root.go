package route

import "net/http"

// RootHandler handles requests for the root address ("/") of the API.
//  GET /
func RootHandler(writer http.ResponseWriter, request *http.Request) {
	payload := map[string]string{"message": "API root! Nothing to see here. Move along..."}
	MakeJSONResponse(writer, http.StatusOK, payload)
}
