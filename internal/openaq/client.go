package openaq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const DateISO = "2006-01-02"

// Client is a http client for https://api.openaq.org
type Client struct{}

func NewClient() *Client {
	return &Client{}
}

type Response struct {
	Results json.RawMessage `json:"results"`
}

// Cities sends request to the `/cities` endpoint and return list of cities in the country
func (c *Client) Cities(country string, v interface{}) error {
	res, err := http.Get(fmt.Sprintf("https://api.openaq.org/v1/cities?country=%s", country))
	if err != nil {
		return fmt.Errorf("http request failed, err %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("body data corrupted, err %v", err)
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("unmarshal failed, err %v", err)
	}

	err = json.Unmarshal(result.Results, v)
	if err != nil {
		return fmt.Errorf("unmarshal failed, err %v", err)
	}

	return nil
}

func (c *Client) Measurements(city string, dateFrom, dateTo time.Time, v interface{}) error {
	res, err := http.Get(fmt.Sprintf("https://api.openaq.org/v1/measurements?city=%s&date_from=%s&date_to=%s", city, dateFrom.Format(DateISO), dateTo.Format(DateISO)))
	if err != nil {
		return fmt.Errorf("http request failed, err %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("body data corrupted, err %v", err)
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("unmarshal failed, err %v", err)
	}

	err = json.Unmarshal(result.Results, v)
	if err != nil {
		return fmt.Errorf("unmarshal failed, err %v", err)
	}

	return nil
}
