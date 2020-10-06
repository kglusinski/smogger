package smogger

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"reflect"
	mock "smogger/testdata/mock"
	"testing"
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
