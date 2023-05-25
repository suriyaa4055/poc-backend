package service

import (
	data "testBench/datalayer"
	"testBench/initializers"
	"testBench/models"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BenchService interface {
	FetchAllHeadUnits() []models.TestBenchTable
	FetchGenData(gen string) ([]models.GenDetails, error)
	FetchInfoByVin(w http.ResponseWriter, r *http.Request)
	FetchInfoByGenVer(w http.ResponseWriter, r *http.Request)
}

type FetchData struct {
}

func GetData() data.DataLayer {
	rout := data.Data{}
	return rout
}

func (d FetchData) FetchAllHeadUnits() []models.TestBenchTable {

	serviceObj := GetData()
	return serviceObj.HeadUnitsFromDB()

}

// Fetch all the vin and other details of a particular HU Generation
func (d FetchData) FetchGenData(gen string) ([]models.GenDetails, error) {

	serviceObj := GetData()

	return serviceObj.GenDataFromDB(gen)

}

func (d FetchData) FetchInfoByVin(w http.ResponseWriter, r *http.Request) {
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

func (d FetchData) FetchInfoByGenVer(w http.ResponseWriter, r *http.Request) {

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
