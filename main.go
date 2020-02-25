package main

import (
	"fmt"
	log "frontline-common/goutils/fllogger"
	"github.com/gorilla/mux"
	"net/http"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/v1"
	"time"
)

// @title Template API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @BasePath /v1
func main() {
	log.Init("GoTemplate")
	log.Info("Starting go-template API Server!")
	router := mux.NewRouter()
	v1.AddRoutes(router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println(srv.ListenAndServe())
}
