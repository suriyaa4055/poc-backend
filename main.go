package main

import (
	"log"
	"net/http"
	"testBench/initializers"

	"github.com/gorilla/mux"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	r := mux.NewRouter()

	print("hi")

	r.HandleFunc("/fetchDetails", rout.FetchInfo).Methods("GET")
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
