package openaq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const DateISO = "2006-01-02"

// Client is a http client for https://api.openaq.org
type Client struct{}

// NewClient creates new openaq API client
func NewClient() *Client {
	return &Client{}
}

// Cities sends request to the `/cities` endpoint and returns list of cities in the country
func (c *Client) Cities(country string, v interface{}) error {
	res, err := http.Get(fmt.Sprintf("https://api.openaq.org/v1/cities?country=%s", country))
	if err != nil {
		return fmt.Errorf("http request failed, err %v", err)
	}
	defer res.Body.Close()

	return process(res, v)
}

// Measurements sends request to the `/measurements` endpoint and returns list of measurements in the city
func (c *Client) Measurements(city string, dateFrom, dateTo time.Time, v interface{}) error {
	res, err := http.Get(fmt.Sprintf("https://api.openaq.org/v1/measurements?city=%s&date_from=%s&date_to=%s&parameter=pm25", url.QueryEscape(city), dateFrom.Format(DateISO), dateTo.Format(DateISO)))
	if err != nil {
		return fmt.Errorf("http request failed, err %v", err)
	}
	defer res.Body.Close()

	return process(res, v)
}

type response struct {
	Results json.RawMessage `json:"results"`
}

func process(res *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("body data corrupted, err %v", err)
	}

	var result response
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
