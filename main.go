package main

import (
	"log"
	"net/http"
	"os"

	dbconnector "./db"
	"github.com/joho/godotenv"
)

func main() {
	// Init ENV variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to database
	dbconnector.ConnectMongoDB()
	// Starting the server
	http.Handle("/", GetRoutes())
	log.Println("Server Started listening", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
