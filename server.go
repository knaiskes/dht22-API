package main

import (
	"fmt"
	"github.com/KNaiskes/measurementsTH-API/api"
	"github.com/KNaiskes/measurementsTH-API/db"
	"net/http"
)

func main() {
	db.InitDB()

	measurementHandlers := api.MakeMeasurementsHandlers()

	http.HandleFunc("/measurements", measurementHandlers.Measurements)
	http.HandleFunc("/measurements/", measurementHandlers.GetMeasurement)

	fmt.Println("Server started at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
