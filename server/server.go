package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//Application struct to hold middleware, router
type Application struct {
	MiddleWare *negroni.Negroni
	Router     *mux.Router
}

//NewWeb Creates a new application for handling of Web Application
func NewWeb(root string, router *mux.Router) *Application {
	//Setting up negroni for MiddleWare
	middleware := negroni.New()

	//Add middleware onto stack
	middleware.Use(negroni.NewRecovery())
	middleware.Use(negroni.NewLogger())
	middleware.Use(negroni.NewStatic(http.Dir("client")))

	//Use gorilla/mux for routing
	if router == nil {
		router = mux.NewRouter()
	}
	middleware.UseHandler(router)

	//returb application struct
	return &Application{middleware, router}
}

//NewAPI Creates a new application for handling of Restful API Application
func NewAPI(root string, router *mux.Router) *Application {
	//Setting up negroni for MiddleWare
	middleware := negroni.New()

	//Add middleware onto stack
	middleware.Use(negroni.NewRecovery())
	middleware.Use(negroni.NewLogger())

	//set up middleware to use provided router
	middleware.UseHandler(router)

	//returb application struct
	return &Application{middleware, router}
}
