package api

import (
	"encoding/json"
	"fmt"
	"github.com/KNaiskes/measurementsTH-API/db"
	"github.com/KNaiskes/measurementsTH-API/models"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type measurementHandlers struct {
	sync.Mutex
}

func (h *measurementHandlers) Measurements(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	}
}

// GET Method

func (h *measurementHandlers) get(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	results := db.GetAll()
	h.Unlock()

	jsonData, err := json.Marshal(results)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// GET a single measurement by id
func (h *measurementHandlers) GetMeasurement(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	measurement, notFound := db.GetOne(parts[2])
	h.Unlock()

	if notFound {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(measurement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// GET measurements by name
func (h *measurementHandlers) GetMeasurementsByName(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	measurement := db.GetAllByName(parts[3])
	h.Unlock()
	/*
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	*/

	jsonData, err := json.Marshal(measurement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// POST
func (h *measurementHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		// TODO: let user know what gone bad
		return
	}

	var measurement models.Measurement
	err = json.Unmarshal(bodyBytes, &measurement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	measurement.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	db.InsertNewMeasurement(measurement)
	defer h.Unlock()

}

func MakeMeasurementsHandlers() *measurementHandlers {
	return &measurementHandlers{}
}
