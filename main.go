package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	http.Handle("/", getRoutes())
	log.Println("Server Started listening", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
