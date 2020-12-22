package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/datastore"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

const (
	srvAddr     = ":5000"
	metricsAddr = ":5001"
)

// @title Template API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @BasePath /
func main() {
	log, flushLog, err := utils.LogInit()
	if err != nil {
		return
	}
	defer flushLog()

	log.Info("opening connection to datastore")
	db, err := datastore.OpenDB()
	if err != nil {
		log.Fatal("OpenDB: %s", err)
		return
	}

	log.Info("migrating datastore")
	err = datastore.Migrate(db)
	if err != nil {
		log.Fatal("Migrate DB: %s", err)
		return
	}

	log.Info("initializing helloworld Service")
	appServer := server.New(log, db)
	router := appServer.Router(
		server.MetricMiddleware(),
	)

	srv := &http.Server{
		Handler: router,
		Addr:    srvAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Serve our handler.
	go func() {
		log.Info("server listening at %s", srvAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("error while serving: %s", err)
		}
	}()

	// Serve our metrics.
	go func() {
		log.Info("metrics listening at %s", metricsAddr)
		if err := http.ListenAndServe(metricsAddr, promhttp.Handler()); err != nil {
			log.Fatal("error while serving metrics: %s", err)
		}
	}()

	// Wait until some signal is captured.
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC
}
