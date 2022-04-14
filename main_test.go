package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	hello(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Hello, World.") {
		t.Errorf(`response body "%s" does not contain "Hello, World."`, wr.Body.String())
	}
}

func TestWeatherapi(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/weather?location=Rangoon", nil)

	weatherapi(wr, req)

	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), `"name":"Yangon"`) {
		t.Errorf(`response body "%s" does not contain "name":"Yangon"`, wr.Body.String())
	}
}
