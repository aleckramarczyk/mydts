package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var err error
	//TODO allow specification of config path
	LoadAppConfig(".")

	err = ConnectDB(GenerateConfigString())
	if err != nil {
		log.Println("Failed to connect to database:")
		log.Fatal(err)
	}

	err = Migrate()
	if err != nil {
		log.Println("Failed to migrate database: ")
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	registerApiRoutes(router)

	log.Printf("Starting server on port %s", AppConfig.API_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.API_port), router))
}
