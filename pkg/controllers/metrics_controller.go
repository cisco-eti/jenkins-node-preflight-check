package controllers

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/metrics"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

func init() {
	prometheus.MustRegister(metrics.DeviceCounter)
	prometheus.MustRegister(metrics.PetFamilyCounter)
	prometheus.MustRegister(metrics.PetTypeCounter)
}

// MetricsController struct
type MetricsController struct {
}

// AddRoutes add metrics routes to the Mux router
func (metricCtrl *MetricsController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/metrics", metricCtrl.Get).Methods("GET")
	return router
}

// Get Handler for prometheus (/metrics endpoint)
func (metricCtrl *MetricsController) Get(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/metrics request received")
	promhttp.Handler().ServeHTTP(w, r)
}
