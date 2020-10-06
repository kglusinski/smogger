package main

import (
	"encoding/json"
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
	sRouter.HandleFunc("/cities", getCities).Methods("GET").Queries("country", "{[A-Z]}")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type ErrResponse struct {
	Error string `json:"error"`
}

func whoami(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _  = w.Write([]byte(`{"name": "Smogger v1"}`))
}

func getCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	country := r.URL.Query().Get("country")

	cities, err := s.GetCities(country)
	if err != nil {
		log.Printf("client error, err: %+v", err)
		errRes, _ := json.Marshal(ErrResponse{
			Error: "Something went wrong",
		})

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errRes)
		return
	}

	resp, err := json.Marshal(cities)
	if err != nil {
		log.Printf("json marshal error, err: %+v", err)
		errRes, _ := json.Marshal(ErrResponse{
			Error: "Something went wrong",
		})

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
