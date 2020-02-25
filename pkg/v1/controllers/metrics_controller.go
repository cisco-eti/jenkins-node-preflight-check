package controllers

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/services"
)

func init() {
	_ = prometheus.Register(services.Counter)
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
	promhttp.Handler().ServeHTTP(w, r)
}
