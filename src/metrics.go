package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	svcWatch = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "svc_watch",
			Help: "The current number of services with cert annotations",
		})

	metricsCertsInProgress = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "certs_in_progress",
			Help: "The current number of certificates in progress. Most likely there is a isue!",
		})
)

func registerMetrics() {
	r := prometheus.NewRegistry()
	r.MustRegister(svcWatch)
	// r.MustRegister(metricsCertsInProgress)

	// http.Handle("/metrics/", promhttp.Handler())
	http.Handle("/metrics/", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
}
