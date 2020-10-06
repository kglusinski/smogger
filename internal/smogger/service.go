package smogger

import (
	"fmt"
	"time"
)

type ApiClient interface {
	Cities(country string, v interface{}) error
	Measurements(city string, dateFrom, dateTo time.Time, v interface{}) error
}

type Service struct {
	client ApiClient
}

func NewService(c ApiClient) *Service {
	return &Service{
		client: c,
	}
}

func (s *Service) GetCities(country string) ([]City, error) {
	var cities []City

	err := s.client.Cities(country, &cities)
	if err != nil {
		return []City{}, fmt.Errorf("couldn't get cities from provider, err: %v", err)
	}

	return cities, err
}

func (s *Service) GetMeasurements(city string, dateFrom, dateTo time.Time) ([]Measurement, error) {
	var measurements []Measurement

	err := s.client.Measurements(city, dateFrom, dateTo, &measurements)
	if err != nil {
		return []Measurement{}, fmt.Errorf("couldn't get measurements from provider, err: %v", err)
	}

	return measurements, err
}
