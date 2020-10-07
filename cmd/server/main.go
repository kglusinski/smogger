package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"smogger/internal/openaq"
	"smogger/internal/smogger"
)

const dateRegex string = "(?:[12]\\d{3}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\\d|3[01]))"

var s *smogger.Service

func main() {
	port := os.Getenv("API_PORT")
	log.Println("Starting server - Smogger v1.0 by Kamil Głusiński")

	c := openaq.NewClient()
	s = smogger.NewService(c)

	handle(port)
}

func handle(port string) {
	router := mux.NewRouter().StrictSlash(true)
	sRouter := router.PathPrefix("/v1").Subrouter()

	sRouter.HandleFunc("/", whoami)
	sRouter.HandleFunc("/cities", getCities).
		Methods("GET").
		Queries("country", "{[A-Z]}")
	sRouter.HandleFunc("/measurements", getMeasurements).
		Methods("GET").
		Queries("city", "{[A-Z]}", "date_from", "{date_from:"+dateRegex+"}", "date_to", "{date_to:"+dateRegex+"}")

	log.Fatal(http.ListenAndServe(":"+ port, router))
}
