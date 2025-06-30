package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    return json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
    type envelope struct {
        Error string `json:"error"`
    }

    return writeJSON(w, status, &envelope{Error: message})
}

func readJSON(w http.ResponseWriter,r *http.Request, data any) error {
    maxBytes := 1_048_576 // 1 MB
    r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() // Prevents unknown fields in JSON

    return decoder.Decode(data)
}
