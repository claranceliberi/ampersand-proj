package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type BatteryTelematics struct {
	BatteryId    string  `json:"battery_id"`
	Longitude    string  `json:"longitude"`
	Latitude     string  `json:"latitude"`
	BatterySoc   float64 `json:"battery_soc"`
	Charging     bool    `json:"charging"`
	ChargingRate float64 `json:"charging_rate"`
}

func sendTelemetryData(wg *sync.WaitGroup, i int) {
	defer wg.Done() // Notify that this goroutine is done after function execution

	data := BatteryTelematics{
		BatteryId:    fmt.Sprintf("%d", rand.Intn(2)+1),
		Longitude:    fmt.Sprintf("%.4f° N", rand.Float64()+1.9355),
		Latitude:     fmt.Sprintf("%.4f° E", rand.Float64()+30.0655),
		BatterySoc:   rand.Float64() * 100,
		Charging:     rand.Intn(2) == 1,
		ChargingRate: rand.Float64() * 10,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error preparing request: %v\n", err)
		return
	}

	resp, err := http.Post("http://localhost:8059/api/v1/battery_telematics", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Response from server for request %d: Status %s, Body %s\n", i, resp.Status, body)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup // WaitGroup to wait until all goroutines are done

	for i := 0; i < 1000; i++ {
		wg.Add(1) // Add 1 to the WaitGroup counter for each new goroutine
		go sendTelemetryData(&wg, i)
		time.Sleep(50 * time.Millisecond) // Add sleep if necessary to avoid overwhelming the server
	}

	wg.Wait() // Block until all goroutines in the WaitGroup are done
}
