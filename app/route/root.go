package route

import "net/http"

// RootHandler handles requests for the root address ("/") of the API.
//  GET /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root root!"))
}
