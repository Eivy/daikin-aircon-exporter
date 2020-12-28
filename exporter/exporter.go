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
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listen, nil))
}
