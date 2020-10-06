package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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

func getMeasurements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	city := r.URL.Query().Get("city")
	dateFrom, _ := time.Parse("2006-01-02", r.URL.Query().Get("date_from"))
	dateTo, _ := time.Parse("2006-01-02", r.URL.Query().Get("date_to"))

	cities, err := s.GetMeasurements(city, dateFrom, dateTo)
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
