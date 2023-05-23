package handlers

import (
	"encoding/json"
	"net/http"
)

type BenchInterface interface {
	FetchInfo(w http.ResponseWriter, r *http.Request)
}

type Database struct {
}

func (d Database) FetchInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("Created")
	w.Header().Set("content-Type", "application/json")

	json.NewEncoder(w).Encode()
}
