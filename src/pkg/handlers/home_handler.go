package handlers

import (
	"net/http"
)

// API Home path handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	retStr := "Demo Home"
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", string(len(retStr)))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(retStr))
	}
}
