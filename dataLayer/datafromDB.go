package datalayer

import (
	"fmt"
	"testBench/initializers"
	"testBench/models"
)

type DataLayer interface {
	HeadUnitsFromDB() []models.TestBenchTable
	GenDataFromDB(gen string) ([]models.Details, error)
}

type Data struct {
}

// All the headunits From database
func (d Data) HeadUnitsFromDB() []models.TestBenchTable {

	var lists []models.TestBenchTable
	if result := initializers.DB.Find(&lists); result.Error != nil {
		fmt.Println(result.Error)
	}

	return lists
	//json.NewEncoder(w).Encode(lists)
}

//Get all the details of particular headunit except services availale
unc (d Data) GenDataFromDB(gen string) ([]models.Details, error) {

	var tb []models.Details
	result := initializers.D.Where("hu_gen = ?", gen).Find(&tb)
	if result.Error != nil {
		
		eturn []models.Details{}, result.Error
}

	return tb, nil
}
