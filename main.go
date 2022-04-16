package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// Add API key from OpenWeatherMap
const key = ""

func NewClient(key string) *Client {
	return &Client{
		APIKey:  key,
		BaseURL: "https://api.openweathermap.org",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c Client) FormatURL(location string) string {
	location = url.QueryEscape(location)
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) ([]byte, int, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return []byte(""), http.StatusInternalServerError, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return []byte(""), http.StatusNotFound, fmt.Errorf("could not find location: %s ", location)
	}
	if resp.StatusCode != http.StatusOK {
		return []byte(""), http.StatusBadRequest, fmt.Errorf("unexpected response status %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), http.StatusInternalServerError, err
	}

	return data, http.StatusOK, nil
}

func GetWResult(location, key string) ([]byte, int, error) {
	c := NewClient(key)
	weather, code, err := c.GetWeather(location)
	if err != nil {
		return []byte(""), code, err
	}
	return weather, code, nil
}

func weatherapi(w http.ResponseWriter, r *http.Request) {
	location := r.FormValue("location")
	w.Header().Set("Content-Type", "application/json")
	data, code, err := GetWResult(location, key)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	w.Write(data)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello, World.")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/weather", weatherapi)
	http.ListenAndServe(":8080", nil)
}
