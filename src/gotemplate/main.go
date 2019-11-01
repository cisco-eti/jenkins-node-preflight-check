package gotemplate

import (
	"fmt"
	"log"
	"net/http"
)

func getDeviceZone(w http.ResponseWriter, r *http.Request) {

	device := r.URL.Path[len("/deviceZone/"):]
	if device == "A" {
		fmt.Fprint(w, "Plumbing")
		return
	}

	if device == "B" {
		fmt.Fprint(w, "Gardening")
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1", getDeviceZone)
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
