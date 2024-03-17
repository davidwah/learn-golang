package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	waktu := time.NewTicker(15 * time.Second)
	defer waktu.Stop()

	go func() {
		for {
			select {
			case <-waktu.C:
				updateStatus()
			}
		}
	}()

	http.HandleFunc("/", statusHandler)
	fmt.Println("Service running on port 8080...")
	fmt.Println("Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// fungsi update status data Water dan Wind
func updateStatus() {
	rand.Seed(time.Now().UnixNano())
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	// struct Status
	status := Status{
		Water: water,
		Wind:  wind,
	}

	statusJSON, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		fmt.Println("Error encoding status to JSON:", err)
		return
	}

	// tulis data format json
	err = ioutil.WriteFile("status.json", statusJSON, 0644)
	if err != nil {
		fmt.Println("Error writing status JSON to file:", err)
		return
	}

	fmt.Println("Status updated:", status)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Read status JSON from file
	statusJSON, err := ioutil.ReadFile("status.json")
	if err != nil {
		fmt.Println("Error reading status JSON from file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Decode status JSON
	var status Status
	err = json.Unmarshal(statusJSON, &status)
	if err != nil {
		fmt.Println("Error decoding status JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Determine status for water
	var waterStatus string
	switch {
	case status.Water < 5:
		waterStatus = "Aman"
	case status.Water >= 5 && status.Water <= 8:
		waterStatus = "Siaga"
	default:
		waterStatus = "Bahaya"
	}

	// Determine status for wind
	var windStatus string
	switch {
	case status.Wind < 6:
		windStatus = "Aman"
	case status.Wind >= 6 && status.Wind <= 15:
		windStatus = "Siaga"
	default:
		windStatus = "Bahaya"
	}

	// Display status
	fmt.Fprintf(w, "Status Air: %d meter (%s)\n", status.Water, waterStatus)
	//fmt.Printf(w, "Status Air: %d meter (%s)\n", status.Water, waterStatus)
	fmt.Fprintf(w, "Status Angin: %d meter/detik (%s)\n", status.Wind, windStatus)
	//fmt.Printf(w, "Status Angin: %d meter/detik (%s)\n", status.Wind, windStatus)
}
