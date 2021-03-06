package main

import (
	"./dao"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"github.com/mhaddon/gke-k8s/services/common/src/helper"
	"log"
	"net/http"
)

func main() {
	conf := config.GetInstance()

	log.Printf("\nHello, serving Airport REST API on port :%v", conf.Http.Port)

	router := mux.NewRouter()

	router.HandleFunc("/health", helper.IsDBHealthy).Methods("GET")
	router.HandleFunc("/alive", helper.IsDBHealthy).Methods("GET")

	router.HandleFunc("/airports", dao.GetAirports).Methods("GET")
	router.HandleFunc("/airports/country_code/{code}", dao.GetAirportsByCountryCode).Methods("GET")
	router.HandleFunc("/airports/country_code/{code}/search/{query}", dao.GetAirportsByCountryCodeAndSearch).Methods("GET")
	router.HandleFunc("/airports/{id}", dao.GetAirport).Methods("GET")
	router.HandleFunc("/airports/{id}", dao.CreateAirport).Methods("POST")
	router.HandleFunc("/airports/{id}", dao.DeleteAirport).Methods("DELETE")
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
