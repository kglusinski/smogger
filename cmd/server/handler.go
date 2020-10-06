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
	_, _ = w.Write([]byte(`{"name": "Smogger v1"}`))
}

func getCities(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")

	writer := NewWriter(w, r)
	writer.writeResponse(s.GetCities(country))
}

func getMeasurements(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	dateFrom, _ := time.Parse("2006-01-02", r.URL.Query().Get("date_from"))
	dateTo, _ := time.Parse("2006-01-02", r.URL.Query().Get("date_to"))

	writer := NewWriter(w, r)
	writer.writeResponse(s.GetMeasurements(city, dateFrom, dateTo))
}

type Writer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewWriter(w http.ResponseWriter, r *http.Request) *Writer {
	return &Writer{
		w: w,
		r: r,
	}
}

func (h *Writer) writeResponse(entity interface{}, err error) {
	h.w.Header().Set("content-type", "application/json")
	h.w.Header().Set("Access-Control-Allow-Origin", "*")

	if h.r.Method == http.MethodOptions {
		return
	}

	if err != nil {
		log.Printf("client error, err: %+v", err)
		errRes, _ := json.Marshal(ErrResponse{
			Error: "Something went wrong",
		})

		h.w.WriteHeader(http.StatusInternalServerError)
		_, _ = h.w.Write(errRes)
		return
	}

	resp, err := json.Marshal(entity)
	if err != nil {
		log.Printf("json marshal error, err: %+v", err)
		errRes, _ := json.Marshal(ErrResponse{
			Error: "Something went wrong",
		})

		h.w.WriteHeader(http.StatusInternalServerError)
		_, _ = h.w.Write(errRes)
		return
	}

	h.w.WriteHeader(http.StatusOK)
	_, _ = h.w.Write(resp)
}
