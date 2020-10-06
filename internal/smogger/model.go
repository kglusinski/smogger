package smogger

import "time"

type City struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Measurement struct {
	Country string `json:"country"`
	City string `json:"city"`
	Location string `json:"location"`
	Parameter string `json:"parameter"`
	Unit string `json:"unit"`
	Value float32 `json:"value"`
	Date struct{
		Utc time.Time `json:"utc"`
		Local time.Time `json:"local"`
	} `json:"date"`
}
