package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

//routes.HandleFunc("/api/books", GetBooks).Methods("GET")

// ControllerFunc type
type ControllerFunc func(http.ResponseWriter, *http.Request)

// IRoute interface
type IRoute interface {
	Get(string, ControllerFunc) IRoute
	Post(string, ControllerFunc) IRoute
	Put(string, ControllerFunc) IRoute
	Delete(string, ControllerFunc) IRoute
	Router() *mux.Router
}

// Route struct
type Route struct{}

// Get command
func (route Route) Get(path string, cf ControllerFunc) IRoute {
	router.HandleFunc(path, cf).Methods("GET")
	return route
}

// Post command
func (route Route) Post(path string, cf ControllerFunc) IRoute {
	router.HandleFunc(path, cf).Methods("POST")
	return route
}

// Put command
func (route Route) Put(path string, cf ControllerFunc) IRoute {
	router.HandleFunc(path, cf).Methods("PUT")
	return route
}

// Delete command
func (route Route) Delete(path string, cf ControllerFunc) IRoute {
	router.HandleFunc(path, cf).Methods("DELETE")
	return route
}

// Router returned
func (route Route) Router() *mux.Router {
	return router
}
