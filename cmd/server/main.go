package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//TODO allow specification of config path
	LoadAppConfig(".")
	ConnectDB(GenerateConfigString())
	Migrate()

	router := mux.NewRouter().StrictSlash(true)
	registerApiRoutes(router)

	log.Printf("Starting server on port %s", AppConfig.API_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.API_port), router))
}
