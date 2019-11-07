package main

import (
	"fmt"
	"log"
	"net/http"

	"sqbu-github.cisco.com/Nyota/go-template/pkg/apis/devicezone"
)

func main() {
	fmt.Println("Starting Sample API Server!")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/deviceZone/A", devicezone.GetDeviceZone)
	mux.HandleFunc("/api/v1/deviceZone/B", devicezone.GetDeviceZone)
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
