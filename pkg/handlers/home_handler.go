package handlers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// API Home path handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	retStr := "Demo Home"
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", string(len(retStr)))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(retStr))
	}
}

var (
	Counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "deviceZoneReq",
			Help: "This is my counter",
		},
		[]string{"device"},
	)
)

// Handler for prometheus (/metrics endpoint)
func PromMetrics(w http.ResponseWriter, r *http.Request) {

	promhttp.Handler().ServeHTTP(w, r)
	prometheus.Register(Counter)
}
