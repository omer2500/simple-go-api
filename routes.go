package main

import (
	"./components/user"
	"github.com/gorilla/mux"
)

func getRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", user.CreateUser).Methods("POST")
	return r
}
