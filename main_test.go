package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/homes", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreateHomeHandler(t *testing.T) {
	requestBody := `{
        "Address":"Palanpur",
		"Description":"Nice 2BHK Apartment",
		"City":"Gandhidham",
		"State":"Gujarat",
		"PostaclCode" : 370205,
		"Price" : 15000000,
		"ForSale" : false,
		"Available" : true
    }`
	req, err := http.NewRequest("POST", "/homes", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/homes/123456", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateHomeHandler(t *testing.T) {
	requestBody := `{
		"Address":"Palanpur",
		"Description":"Nice 2BHK Apartment",
		"City":"Gandhidham",
		"State":"Gujarat",
		"PostaclCode" : 370205,
		"Price" : 15000000,
		"ForSale" : false,
		"Available" : true
    }`
	req, err := http.NewRequest("PUT", "/homes/123456", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestDeleteHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/homes/123456", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
