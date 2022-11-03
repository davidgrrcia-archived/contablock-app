package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Failure(w http.ResponseWriter, err error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
}

func Success(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(i)
}
