package routes

import "net/http"

var routeString = "/test"

// HomeHandler The route and handler used for homepage
type HomeHandler struct {
	//Route
}

//GetRoute gets the roo
func (h HomeHandler) GetRoute() string {
	return routeString
}

//Handle Home routing
func (h HomeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test"))
}
