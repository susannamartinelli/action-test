package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestStatusHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/status", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(statusHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := map[string]string{"version": version}
    var actual map[string]string
    if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
        t.Fatalf("could not decode json response: %v", err)
    }

    if actual["version"] != expected["version"] {
        t.Errorf("handler returned unexpected body: got %v want %v",
            actual["version"], expected["version"])
    }
}
