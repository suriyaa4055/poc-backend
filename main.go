package main

import (
	"log"
	"net/http"
	"testBench/handlers"
	"testBench/initializers"

	"github.com/gorilla/mux"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	r := mux.NewRouter()
	rout := handlers.Database{}

	print("hi")

	r.HandleFunc("/fetchDetails", rout.FetchInfo).Methods("GET")
	r.HandleFunc("/fetchDetails/{HuGen}", rout.FetchInfoByGen).Methods("GET")
	r.HandleFunc("/fetchDetails/{HuGen}/{HuVer}", rout.FetchInfoByGenVer).Methods("GET")
	r.HandleFunc("/fetchAllDetails/{vin}", rout.FetchInfoByVin).Methods("GET")

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
