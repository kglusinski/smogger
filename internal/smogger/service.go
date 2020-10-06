package smogger

import "fmt"

type ApiClient interface {
	Cities(country string, v interface{}) error
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
