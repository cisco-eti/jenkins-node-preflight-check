package main

import (
    "log"
    "net/http"
    "fmt"
)

func DeviceZoneServer(w http.ResponseWriter, r *http.Request) {

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
    mux.HandleFunc("/api/v1", DeviceZoneServer)
    if err := http.ListenAndServe(":5000", mux); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
