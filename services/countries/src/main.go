package main

import (
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"./dao"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	conf := config.GetInstance()

	router := mux.NewRouter()
	fmt.Printf("\nHello, serving on :%v", conf.Http.Port)
	router.HandleFunc("/countries", dao.GetCountries).Methods("GET")
	router.HandleFunc("/countries/{id}", dao.GetCountry).Methods("GET")
	router.HandleFunc("/countries/{id}", dao.CreateCountry).Methods("POST")
	router.HandleFunc("/countries/{id}", dao.DeleteCountry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}