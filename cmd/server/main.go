package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"smogger/internal/openaq"
	"smogger/internal/smogger"
)

var s *smogger.Service

func main() {
	log.Println("Starting server - Smogger v1.0 by Kamil Głusiński")

	c := openaq.NewClient()
	s = smogger.NewService(c)

	handle()
}

func handle() {
	router := mux.NewRouter().StrictSlash(true)
	sRouter := router.PathPrefix("/v1").Subrouter()

	sRouter.HandleFunc("/", whoami)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func whoami(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name": "Smogger v1"}`))
}
