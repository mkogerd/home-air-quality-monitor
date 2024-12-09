package collector

import (
	"awair-exporter/awair"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type AirQualityCollector struct {
	apiURL              string
	score               *prometheus.GaugeVec
	dewPoint            *prometheus.GaugeVec
	temperature         *prometheus.GaugeVec
	relativeHumidity    *prometheus.GaugeVec
	absoluteHumidity    *prometheus.GaugeVec
	co2                 *prometheus.GaugeVec
	co2Estimate         *prometheus.GaugeVec
	co2EstimateBaseline *prometheus.GaugeVec
	voc                 *prometheus.GaugeVec
	vocBaseline         *prometheus.GaugeVec
	vocHydrogenRaw      *prometheus.GaugeVec
	vocEthanolRaw       *prometheus.GaugeVec
	pm25                *prometheus.GaugeVec
	pm10Estimate        *prometheus.GaugeVec
}

// NewAirQualityCollector creates a new AirQualityCollector
func NewAirQualityCollector(apiURL string) *AirQualityCollector {
	return &AirQualityCollector{
		apiURL: apiURL,
		score: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_score",
				Help: "Awair Score (0-100), overall summary of air quality",
			},
			[]string{"location"},
		),
		dewPoint: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_dew_point_celsius",
				Help: "The temperature at which water will condense and form into dew",
			},
			[]string{"location"},
		),
		relativeHumidity: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_humidity_relative",
				Help: "Current relative humidity percentage, describes how saturated the air is with respect to water at a given temperature",
			},
			[]string{"location"},
		),
		absoluteHumidity: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_humidity_absolute_grams_per_cubic_meter",
				Help: "The amount of water vapor in the air (g/m³)",
			},
			[]string{"location"},
		),
		temperature: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_temperature_celsius",
				Help: "Current temperature in Celsius",
			},
			[]string{"location"},
		),
		co2: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_co2_ppm",
				Help: "Current CO2 levels in parts per million",
			},
			[]string{"location"},
		),
		co2Estimate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_co2_estimate_ppm",
				Help: "Current estimated CO2 levels in parts per million, calculated by the TVOC sensor",
			},
			[]string{"location"},
		),
		co2EstimateBaseline: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_co2_estimate_baseline",
				Help: "A unitless value that represents the baseline from which the TVOC sensor partially derives its estimated CO2 output",
			},
			[]string{"location"},
		),
		voc: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_voc_ppb",
				Help: "Total Volatile Organic Compounds (ppb)",
			},
			[]string{"location"},
		),
		vocBaseline: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_voc_baseline",
				Help: "A unitless value that represents the baseline from which the TVOC sensor partially derives its TVOC output",
			},
			[]string{"location"},
		),
		vocHydrogenRaw: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_voc_hydrogen_raw",
				Help: "A unitless value that represents the Hydrogen gas signal from which the TVOC sensor partially derives its TVOC output",
			},
			[]string{"location"},
		),
		vocEthanolRaw: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_voc_ethanol_raw",
				Help: "A unitless value that represents the Ethanol gas signal from which the TVOC sensor partially derives its TVOC output",
			},
			[]string{"location"},
		),
		pm25: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_pm25_micrograms_per_cubic_meter",
				Help: "Concentration of particulate matter less than 2.5 microns in diameter (µg/m³)",
			},
			[]string{"location"},
		),
		pm10Estimate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "awair_pm10_micrograms_per_cubic_meter",
				Help: "Estimated concentration of particulate matter less than 10 microns in diameter (µg/m³), calculated by the PM2.5 sensor",
			},
			[]string{"location"},
		),
	}
}

// CollectMetrics fetches data from the API and updates metrics
func (c *AirQualityCollector) CollectMetrics() {
	airData, err := awair.ReadAirData(c.apiURL)
	if err != nil {
		log.Printf("Error fetching data from API: %v", err)
		return
	}

	location := "default" // You can adjust this for multiple sensors or locations
	c.score.WithLabelValues(location).Set(float64(airData.Score))
	c.dewPoint.WithLabelValues(location).Set(float64(airData.DewPoint))
	c.temperature.WithLabelValues(location).Set(airData.Temp)
	c.relativeHumidity.WithLabelValues(location).Set(airData.Humid)
	c.absoluteHumidity.WithLabelValues(location).Set(airData.AbsHumid)
	c.co2.WithLabelValues(location).Set(float64(airData.CO2))
	c.co2Estimate.WithLabelValues(location).Set(float64(airData.CO2Est))
	c.co2EstimateBaseline.WithLabelValues(location).Set(float64(airData.CO2EstBaseline))
	c.voc.WithLabelValues(location).Set(float64(airData.VOC))
	c.vocBaseline.WithLabelValues(location).Set(float64(airData.VOCBaseline))
	c.vocHydrogenRaw.WithLabelValues(location).Set(float64(airData.VOCH2Raw))
	c.vocEthanolRaw.WithLabelValues(location).Set(float64(airData.VOCEthanolRaw))
	c.pm25.WithLabelValues(location).Set(float64(airData.PM25))
	c.pm10Estimate.WithLabelValues(location).Set(float64(airData.PM10Est))
}

// Describe sends metric descriptions to Prometheus
func (c *AirQualityCollector) Describe(ch chan<- *prometheus.Desc) {
	c.score.Describe(ch)
	c.dewPoint.Describe(ch)
	c.temperature.Describe(ch)
	c.relativeHumidity.Describe(ch)
	c.absoluteHumidity.Describe(ch)
	c.co2.Describe(ch)
	c.co2Estimate.Describe(ch)
	c.co2EstimateBaseline.Describe(ch)
	c.voc.Describe(ch)
	c.vocBaseline.Describe(ch)
	c.vocHydrogenRaw.Describe(ch)
	c.vocEthanolRaw.Describe(ch)
	c.pm25.Describe(ch)
	c.pm10Estimate.Describe(ch)
}

// Collect sends the metric values to Prometheus
func (c *AirQualityCollector) Collect(ch chan<- prometheus.Metric) {
	log.Print("Collecting data with `Collect`")
	c.score.Collect(ch)
	c.dewPoint.Collect(ch)
	c.temperature.Collect(ch)
	c.relativeHumidity.Collect(ch)
	c.absoluteHumidity.Collect(ch)
	c.co2.Collect(ch)
	c.co2Estimate.Collect(ch)
	c.co2EstimateBaseline.Collect(ch)
	c.voc.Collect(ch)
	c.vocBaseline.Collect(ch)
	c.vocHydrogenRaw.Collect(ch)
	c.vocEthanolRaw.Collect(ch)
	c.pm25.Collect(ch)
	c.pm10Estimate.Collect(ch)
}
