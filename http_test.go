package main

import (
	"net/http"
	"net/url"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
	"github.com/icrowley/fake"
)

func TestNearbyHandlerWithoutParams(t *testing.T){
	req, err := http.NewRequest("GET", "/nearby", nil)
    if err != nil { t.Fatal(err) }
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NearbyHandler)
	handler.ServeHTTP(rr, req)
	expectStatus(t, rr.Code, http.StatusBadRequest)
}

func TestNearbyHandlerWithParams(t *testing.T){
	longitude := fake.Longitude()
	latitude := fake.Latitude()
	url := fmt.Sprintf("/nearby?longitude=%f&latitude=%f", longitude, latitude)
	req, err := http.NewRequest("GET", url, nil)
    if err != nil { t.Fatal(err) }
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NearbyHandler)
	handler.ServeHTTP(rr, req)
	expectStatus(t, rr.Code, http.StatusOK)
}

func TestReservationWithParams(t *testing.T){
	uri := "/reserve"

	data := url.Values{}
    data.Set("name", "foo")
	data.Add("id_restaurant", "1")
	data.Add("date", "2011-01-02T22:04:05Z")

	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    if err != nil { t.Fatal(err) }
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReserveHandler)
	handler.ServeHTTP(rr, req)
	expectStatus(t, rr.Code, http.StatusOK)
}

func TestReservationWithInvalidParams(t *testing.T){
	uri := "/reserve"

	data := url.Values{}
    data.Set("name", "foo")
	data.Add("id_restaurant", "1")
	data.Add("date", "")

	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    if err != nil { t.Fatal(err) }
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReserveHandler)
	handler.ServeHTTP(rr, req)
	expectStatus(t, rr.Code, http.StatusBadRequest)
}

func expectStatus(t *testing.T, code int, status int){
	if code != status {
        t.Errorf("handler returned wrong status code: got %v want %v",
            code, status)
    }
}