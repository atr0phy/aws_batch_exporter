package awsBatchExporter

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Run() {
	addr := getenv("SERVER_ADDR", ":8080")
	region := os.Getenv("REGION")
	if len(region) == 0 {
		log.Fatalln("region is required value. exit an application.")
	}
	collector, err := New(region)
	if err != nil {
		log.Fatalln(err)
	}
	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(addr, nil))
}