package controllers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/metrics"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

func init() {
	prometheus.MustRegister(metrics.DeviceCounter)
	prometheus.MustRegister(metrics.PetFamilyCounter)
	prometheus.MustRegister(metrics.PetTypeCounter)
}

// Get Handler for prometheus (/metrics endpoint)
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/metrics request received")

	promhttp.Handler().ServeHTTP(w, r)
}
