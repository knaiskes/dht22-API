package main

import (
	"github.com/KNaiskes/measurementsTH-API/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMeasurements(t *testing.T) {
	req, err := http.NewRequest("GET", "/measurements", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MakeMeasurementsHandlers().Measurements)

	handler.ServeHTTP(newRecorder, req)

	if status := newRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	/*
		expected := `[{"ID":"id1","Name":"testName","Temperature":"21.00","Humidity":"54.22"},{"ID":"id2","Name":"testName2","Temperature":"30.32","Humidity":"60.22"}]`

		if newRecorder.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				newRecorder.Body.String(), expected)
		}
	*/
}
