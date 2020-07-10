package main

import (
	"encoding/json"
	//	"fmt"
	"net/http"
	"sync"
)

type Measurement struct {
	ID          string `json:id`
	Name        string `json:name`
	Temperature string `json:temperature`
	Humidity    string `json:humidity`
}

//TODO: add a real database

type measurementHandlers struct {
	sync.Mutex
	fakeDB map[string]Measurement
}

func (h *measurementHandlers) measurements(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	}
}

// GET Method

func (h *measurementHandlers) get(w http.ResponseWriter, r *http.Request) {
	measurements := make([]Measurement, len(h.fakeDB))

	h.Lock()

	i := 0
	for _, m := range h.fakeDB {
		measurements[i] = m
		i++
	}

	h.Unlock()

	jsonData, err := json.Marshal(measurements)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func makeMeasurementsHandlers() *measurementHandlers {
	return &measurementHandlers{
		fakeDB: map[string]Measurement{
			"id1": Measurement{
				ID:          "id1",
				Name:        "testName",
				Temperature: "21.00",
				Humidity:    "54.22",
			},
		},
	}
}

func main() {
	measurementHandlers := makeMeasurementsHandlers()

	http.HandleFunc("/measurements", measurementHandlers.measurements)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	//	fmt.Println("Server started")
}
