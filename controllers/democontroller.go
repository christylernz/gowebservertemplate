package controllers

import "net/http"

//Demo demostration of controller function
func Demo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is a demo"))
}
