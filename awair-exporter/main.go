package main

import (
	"awair-exporter/collector"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	awairHost := os.Getenv("AWAIR_HOST")
	c := collector.NewAirQualityCollector(awairHost)

	// Register the custom collector with Prometheus
	prometheus.MustRegister(c)

	// Periodically collect metrics
	go func() {
		for {
			c.CollectMetrics()
			time.Sleep(15 * time.Second) // Scrape every 15 seconds
		}
	}()

	// Expose metrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Serving metrics at :8080/metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
