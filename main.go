package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	c "testBench/controllers"
	"testBench/initializers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	rout := c.Database{}

	print("hi")

	r.HandleFunc("/fetchDetails", rout.FetchInfo).Methods("GET")
	r.HandleFunc("/fetchDetails/{HuGen}", rout.FetchInfoByGen).Methods("GET")
	r.HandleFunc("/fetchDetails/{HuGen}/{HuVer}", rout.FetchInfoByGenVer).Methods("GET")
	r.HandleFunc("/fetchAllDetails/{vin}", rout.FetchInfoByVin).Methods("GET")

	log.Println("Starting server on port " + os.Getenv("APP_PORT") + "...")
	err = http.ListenAndServe(os.Getenv("APP_PORT"), r)

	if err != nil {
		log.Fatal(err)
	}

}
