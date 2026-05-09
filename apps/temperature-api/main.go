package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func main() {
	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		if location == "" {
			http.Error(w, "missing location parameter", http.StatusBadRequest)
			return
		}
		sendTemperatureResponse(w, location)
	})

	http.HandleFunc("/temperature/", func(w http.ResponseWriter, r *http.Request) {
		location := strings.TrimPrefix(r.URL.Path, "/temperature/")
		if location == "" {
			http.Error(w, "missing location parameter", http.StatusBadRequest)
			return
		}
		sendTemperatureResponse(w, location)
	})

	fmt.Println("Temperature API server starting on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func sendTemperatureResponse(w http.ResponseWriter, location string) {
	temperature := rand.Float64()*30 - 10

	response := TemperatureResponse{
		Value:       temperature,
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "normal",
		SensorID:    fmt.Sprintf("sensor-%s", location),
		SensorType:  "temperature",
		Description: fmt.Sprintf("Temperature reading for %s", location),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
