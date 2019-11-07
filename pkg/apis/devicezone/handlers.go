package devicezone

import (
	"net/http"
	"strings"
)

func GetDeviceZone(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/api/v1/deviceZone/A" &&
		r.URL.Path != "/api/v1/deviceZone/B" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	zoneA := []byte("Plumbing")
	zoneB := []byte("Gardening")

	if r.Method == "GET" {

		device := strings.TrimPrefix(r.URL.Path, "/api/v1/deviceZone/")

		if device == "A" {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", "8")
			w.WriteHeader(http.StatusOK)
			w.Write(zoneA)
		}
		if device == "B" {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", "9")
			w.WriteHeader(http.StatusOK)
			w.Write(zoneB)
		}
	}
}
