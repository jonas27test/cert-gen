package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metricsCertsReady = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "certs_ready",
			Help: "The current number of fulfilled certificates and thus secrets created via cert-manager (if with owner ref).",
		})

	metricsCertsInProgress = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "certs_in_progress",
			Help: "The current number of certificates in progress. Most likely there is a isue!",
		})
)

func registerMetrics() {
	r := prometheus.NewRegistry()
	r.MustRegister(metricsCertsReady)
	r.MustRegister(metricsCertsInProgress)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/metrics/", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
}
