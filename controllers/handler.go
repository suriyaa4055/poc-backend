package handlers

import (
	"fmt"
	"testBench/initializers"
	"testBench/models"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BenchInterface interface {
	FetchInfo(w http.ResponseWriter, r *http.Request)
}

type Database struct {
}

func (d Database) FetchInfo(w http.ResponseWriter, r *http.Request) {

	var lists []models.TestBenchTable
	if result := initializers.DB.Find(&lists); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("Created")
	w.Header().Set("content-Type", "application/json")

	json.NewEncoder(w).Encode(lists)
}
func (d Database) FetchInfoByGen(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gen := (vars["HuGen"])
	// if !err {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	var tb []models.GenDetails
	if result := initializers.DB.Where("hu_gen = ?", gen).Find(&tb); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tb)

}
func (d Database) FetchInfoByVin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := (vars["vin"])
	var tb models.Details
	if result := initializers.DB.Where("vin = ?", vin).Find(&tb); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tb)
}

func (d Database) FetchInfoByGenVer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gen := (vars["HuGen"])
	ver := vars["HuVer"]

	// if !err {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	var tb []models.TestBenchTable
	if result := initializers.DB.Where("hu_gen = ? AND hu_version = ?", gen, ver).Find(&tb); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tb)

}