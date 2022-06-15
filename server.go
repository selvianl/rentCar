package main

import (
	"log"
	"net/http"
	"rentCar/db"
	"rentCar/handlers"

	"github.com/gorilla/mux"
)

func main() {
	DB:= db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	// Customer APIs
    router.HandleFunc("/customer", h.GetAllCustomers).Methods(http.MethodGet)
    router.HandleFunc("/customer", h.CreateCustomer).Methods(http.MethodPost)
    router.HandleFunc("/customer/{id}", h.GetCustomer).Methods(http.MethodGet)

	// Vendor APIs
	log.Println("API is running!")
    router.HandleFunc("/vendor", h.GetAllVendors).Methods(http.MethodGet)
    router.HandleFunc("/vendor", h.CreateVendor).Methods(http.MethodPost)
    router.HandleFunc("/vendor/{id}", h.GetVendor).Methods(http.MethodGet)

	// Location APIs
    router.HandleFunc("/location", h.GetAllLocations).Methods(http.MethodGet)
    router.HandleFunc("/location", h.CreateLocation).Methods(http.MethodPost)
    router.HandleFunc("/location/{id}", h.GetLocation).Methods(http.MethodGet)

    http.ListenAndServe(":4000", router)

}