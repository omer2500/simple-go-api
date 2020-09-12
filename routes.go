package main

import (
	"./components/user"
	"github.com/gorilla/mux"
)

// GetRoutes - init the routers for the api
func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", user.CreateUser).Methods("POST")
	return router
}
