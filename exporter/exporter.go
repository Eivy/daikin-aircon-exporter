package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run do run
func Run(listen, ip string) {
	c := Metrics{
		Target: ip,
	}
	prometheus.MustRegister(c)
	log.Printf("Daikin Aircon Exporter version: %s start, target: %s", Version, ip)
	log.Printf("Listeing %s\n", listen)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listen, nil))
}
