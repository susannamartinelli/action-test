package main

import (
    "encoding/json"
    "net/http"
    "os"
)

var version = "1.0.0" // Set this to the current version of your app

func statusHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"version": version}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/status", statusHandler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}
