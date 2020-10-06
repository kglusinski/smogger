package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server - Smogger v1.0 by Kamil Głusiński")

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