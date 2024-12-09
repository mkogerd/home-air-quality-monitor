package awair

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type AirData struct {
	Timestamp      time.Time `json:"timestamp"`        // Time of reading
	Score          int       `json:"score"`            // Integer score
	DewPoint       float64   `json:"dew_point"`        // Dew point in degrees
	Temp           float64   `json:"temp"`             // Temperature
	Humid          float64   `json:"humid"`            // Humidity percentage
	AbsHumid       float64   `json:"abs_humid"`        // Absolute humidity
	CO2            int       `json:"co2"`              // CO2 value
	CO2Est         int       `json:"co2_est"`          // Estimated CO2 value
	CO2EstBaseline int       `json:"co2_est_baseline"` // CO2 estimate baseline
	VOC            int       `json:"voc"`              // VOC level
	VOCBaseline    int       `json:"voc_baseline"`     // VOC baseline
	VOCH2Raw       int       `json:"voc_h2_raw"`       // Raw H2 VOC level
	VOCEthanolRaw  int       `json:"voc_ethanol_raw"`  // Raw ethanol VOC level
	PM25           int       `json:"pm25"`             // Particulate matter 2.5
	PM10Est        int       `json:"pm10_est"`         // Estimated PM10
}

func ReadAirData(awairHost string) (*AirData, error) {
	// Define the URL
	url := fmt.Sprintf("http://%s/air-data/latest", awairHost)

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer response.Body.Close() // Close the response body at the end

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var airData AirData
	err = json.Unmarshal(body, &airData)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
		return nil, err
	}

	// Print the response
	fmt.Println("Response:", string(body))

	return &airData, nil
}
