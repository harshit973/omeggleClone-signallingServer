package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"omeggleClone-signallingServer/databases"
	"omeggleClone-signallingServer/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	databases.ConnectMongo()

	r := mux.NewRouter()
	r.Handle("*", Router.Routes())

	port := "8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
