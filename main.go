package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"sqbu-github.cisco.com/Nyota/go-template/src/pkg/middleware"
	"sqbu-github.cisco.com/Nyota/go-template/src/pkg/handlers"
)

func main() {
	fmt.Println("Starting go-template API Server!")
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/v1/deviceZone/{deviceId}", handlers.DeviceZoneHandler).Methods("GET")
	http.Handle("/", router)
	router.Use(middleware.OAuthMiddleware)

	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
