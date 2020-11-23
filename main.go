package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
	v1 "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1"
	etilog "wwwin-github.cisco.com/eti/sre-go-logger"
)

var logger *etilog.Logger

// @title Template API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @BasePath /v1
func main() {
	logger = utils.LogInit()
	logger.Info("Initializing Hello-world Service")
	router := mux.NewRouter()
	v1.AddRoutes(router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("%v", srv.ListenAndServe())
}
