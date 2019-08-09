package main

import (
	"github.com/christylernz/gowebservertemplate/controllers"
	"github.com/christylernz/gowebservertemplate/server"
	"github.com/gorilla/mux"
)

func main() {
	webrouter := mux.NewRouter()
	webapp := server.NewWeb("/", webrouter)
	//apirouter := webrouter.PathPrefix("/api").Subrouter()
	//restapp := server.NewAPI("/api", apirouter)
	webapp.Register("/demo", controllers.Demo)
	webapp.MiddleWare.Run(":80")

}
