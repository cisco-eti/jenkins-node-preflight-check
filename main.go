package main

import (
	"fmt"
	"net/http"
	"time"

	log "frontline-common/goutils/fllogger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/handlers"
)

func main() {
	log.Init("GoTemplate")
	log.Info("Starting go-template API Server!")
	router := handlers.Router()
	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println(srv.ListenAndServe())
}
