package services

import "github.com/prometheus/client_golang/prometheus"

var (
	// Counter counts operations of a specified type
	Counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "eti_apps_helloworld_api_counter",
			Help: "This is my counter",
		},
		[]string{"device"},
	)
)
