package router

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
    Details string `json:"details"`
}

func RespondWithError(
	w http.ResponseWriter,
	code int,
	msg string,
	err error,
) {
	RespondWithJSON(w, code, ErrorResponse{Message: msg, Details: err.Error()})
}

func RespondWithJSON(
	w http.ResponseWriter,
	code int,
	payload any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

    err := json.NewEncoder(w).Encode(payload)
    if err != nil {
        log.Printf("Error encoding JSON response: %v", err)
    }
}
