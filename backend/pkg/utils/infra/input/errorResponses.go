package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Failure(w http.ResponseWriter, error error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(errorResponse{Error: error.Error()})
	if err != nil {
		return
	}
}

func Success(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		return
	}
}
