package routes

import (
	"net/http"

	"github.com/christylernz/gowebservertemplate/server"
)

//RouteHandler interface for all methods required to handle a route
type RouteHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
	GetRoute() string
}

//Route struct to store information for route
type Route struct {
	RouteHandler
	RouteString string
}

//Register route to an application's router for handling
func Register(app *server.Application, r RouteHandler) {
	app.Router.HandleFunc(r.GetRoute(), r.Handle)
}
