package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"sqbu-github.cisco.com/Nyota/go-template/pkg/handlers"
)

func main() {
	fmt.Println("Starting go-template API Server!")
        router  := handlers.Router()
	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
