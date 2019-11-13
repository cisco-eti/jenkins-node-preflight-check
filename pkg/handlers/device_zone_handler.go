package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeviceZoneHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if r.Method == "GET" {
		deviceId := vars["deviceId"]
		log.Println("DeviceZoneHandler deviceId:" + deviceId)
		if deviceId == "A" {
			zoneA := "Plumbing"
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", string(len(zoneA)))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(zoneA))
		} else if deviceId == "B" {
			zoneB := "Gardening"
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", string(len(zoneB)))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(zoneB))
		} else if deviceId == "C" {
			zoneC := "Lighting"
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", string(len(zoneC)))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(zoneC))
		} else {
			notFound := "Unknown Device Specified"
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", string(len(notFound)))
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(notFound))
		}
	}
}
