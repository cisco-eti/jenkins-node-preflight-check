package main

import (
	"github.com/gorilla/mux"
	"net/http"
	log "sqbu-github.cisco.com/Nyota/frontline-go-logger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/v1"
	"time"
)

// @title Template API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @BasePath /v1
func main() {
	log.LogInitGlobal("nyota-go-template", log.DefaultProdConfig)
	log.Tracer.Infof("Starting go-template API Server!")
	router := mux.NewRouter()
	v1.AddRoutes(router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Tracer.Infof("%v", srv.ListenAndServe())
}
