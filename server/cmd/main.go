package main

import (
	"aleckramarczyk/mydts/server/api"
	"aleckramarczyk/mydts/server/db"
	"aleckramarczyk/mydts/server/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var err error

	//TODO allow specification of config path
	utils.LoadAppConfig(".")

	// Connect to and migrate database
	err = db.ConnectDB(db.GenerateConfigString())
	if err != nil {
		log.Println("Failed to connect to database:")
		log.Fatal(err)
	}

	err = db.Migrate()
	if err != nil {
		log.Println("Failed to migrate database: ")
		log.Fatal(err)
	}

	// Register API endpoints, start http server.
	router := mux.NewRouter().StrictSlash(true)
	api.RegisterApiRoutes(router)

	log.Printf("Starting server on port %s", utils.AppConfig.API_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", utils.AppConfig.API_port), router))
}
