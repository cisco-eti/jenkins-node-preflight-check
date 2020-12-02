package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/datastore"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
	v1 "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1"
	etilog "wwwin-github.cisco.com/eti/sre-go-logger"
)

var logger *etilog.Logger

const (
	srvAddr     = ":5000"
	metricsAddr = ":5001"
)

// @title Template API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @BasePath /v1
func main() {
	logger = utils.LogInit()
	logger.Info("Initializing Hello-world Service")

	// Create http metrics middleware.
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	router := mux.NewRouter()
	router.Use(std.HandlerProvider("", mdlw))
	v1.AddRoutes(router)
	root.AddRoutes(router)
	datastore.MigrateDB()
	srv := &http.Server{
		Handler: router,
		Addr:    srvAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Serve our handler.
	go func() {
		logger.Info("server listening at %s", srvAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Panicf("error while serving: %s", err)
		}
	}()

	// Serve our metrics.
	// Serve our metrics.
	go func() {
		logger.Info("metrics listening at %s", metricsAddr)
		if err := http.ListenAndServe(metricsAddr, promhttp.Handler()); err != nil {
			log.Panicf("error while serving metrics: %s", err)
		}
	}()

	// Wait until some signal is captured.
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC
}
