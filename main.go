package main

import (
	"github.com/christylernz/gowebservertemplate/routes"
	"github.com/christylernz/gowebservertemplate/server"
	"github.com/gorilla/mux"
)

func main() {
	webrouter := mux.NewRouter()
	webapp := server.NewWeb("/", webrouter)
	//apirouter := webrouter.PathPrefix("/api").Subrouter()
	//restapp := server.NewAPI("/api", apirouter)
	home := routes.HomeHandler{} //routes.Route{RouteHandler: &routes.HomeHandler{}, RouteString: "/test"}
	routes.Register(webapp, home)
	webapp.MiddleWare.Run(":80")

}
