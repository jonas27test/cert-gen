package main

import (
	"log"
	"net/http"
)

// EndpointServer starts the server for the healthz and metrics endpoint
func EndpointServer() {
	registerHealthz()
	registerMetrics()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
