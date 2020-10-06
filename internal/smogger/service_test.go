package smogger

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"reflect"
	mock "smogger/testdata/mock"
	"testing"
	"time"
)

func TestService_GetCities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var cities []City

	c := mock.NewMockApiClient(ctrl)
	c.EXPECT().Cities("PL", &cities).Times(1).Do(func(country string, dest interface{}) {
		_ = json.Unmarshal([]byte(`[{"country": "PL", "name": "Cracow"}, {"country": "PL", "name":"Warsaw"}]`), dest)
	})

	expected := []City{
		{
			Country: "PL",
			Name:    "Cracow",
		},
		{
			Country: "PL",
			Name:    "Warsaw",
		},
	}

	s := &Service{
		client: c,
	}

	t.Run("It should returns two cities in Poland", func(t *testing.T) {
		got, err := s.GetCities("PL")
		if err != nil {
			t.Errorf("unexpected error occured, err: %v", err)
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("slices are not equal, got %+v, expected %+v", got, expected)
		}
	})
}

func TestService_GetMeasurements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var measurements []Measurement

	c := mock.NewMockApiClient(ctrl)

	from, _ := time.Parse("2006-01-02", "2020-01-01")
	to, _ := time.Parse("2006-01-02", "2020-10-01")
	c.EXPECT().Measurements("Warszawa", from , to, &measurements).Times(1).Do(func(city string, from, to time.Time, dest interface{}) {
		_ = json.Unmarshal([]byte(`[
{
        "country": "PL",
        "city": "Warszawa",
        "location": "Warszawa-Chrościckiego",
        "parameter": "pm10",
        "unit": "µg/m³",
        "value": 43.04,
        "date": {
            "utc": "2020-10-01T00:00:00Z",
            "local": "2020-10-01T02:00:00+02:00"
        }
    }
]`), dest)
	})

	expected := []Measurement{
		{
			Country: "PL",
			City:   "Warszawa",
			Location: "Warszawa-Chrościckiego",
			Parameter: "pm10",
			Unit: "µg/m³",
			Value: 43.04,
			Date: struct {
				Utc   time.Time `json:"utc"`
				Local time.Time `json:"local"`
			}{Utc: to, Local: to.Local()},
		},
	}

	s := &Service{
		client: c,
	}

	t.Run("It should returns measurement in Warsaw", func(t *testing.T) {
		got, err := s.GetMeasurements("Warszawa", from , to)
		if err != nil {
			t.Errorf("unexpected error occured, err: %v", err)
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("slices are not equal, got %+v, expected %+v", got, expected)
		}
	})
}
