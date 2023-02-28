package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func Failure(w http.ResponseWriter, error error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	var errorRes errorResponse
	errorRes.StatusCode = http.StatusInternalServerError
	errorRes.Error = error.Error()

	err := json.NewEncoder(w).Encode(errorRes)
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
