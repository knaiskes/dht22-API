package main

import (
	//"encoding/json"
	"fmt"
	"github.com/KNaiskes/measurementsTH-API/api"
	"net/http"
	//"strings"
	//"sync"
)

func main() {
	measurementHandlers := api.MakeMeasurementsHandlers()

	http.HandleFunc("/measurements", measurementHandlers.Measurements)
	http.HandleFunc("/measurements/", measurementHandlers.GetMeasurement)

	fmt.Println("Server started at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
