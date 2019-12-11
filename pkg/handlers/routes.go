package handlers

import (

        "github.com/gorilla/mux"
        "sqbu-github.cisco.com/Nyota/go-template/pkg/middleware"
)


func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/v1/deviceZone/{deviceId}", DeviceZoneHandler).Methods("GET")
	router.HandleFunc("/metrics", PromMetrics).Methods("GET")
	router.Use(middleware.OAuthMiddleware)
	return router
}
