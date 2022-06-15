package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentCar/models"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
    Id int `json:"id" gorm:"primaryKey"` 
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    TCKN string `json:"tckn" validate:"required,tckn"`
    BirthDate string `json:"birthDate"`
    PhoneNumber string `json:"phoneNumber"`
}

func ResponseCustomer(customerDB models.Customer) Customer{
	return Customer{
		Id: customerDB.Id,
		FirstName: customerDB.FirstName,
		LastName: customerDB.LastName,
		TCKN: customerDB.TCKN,
		BirthDate: customerDB.BirthDate,
		PhoneNumber: customerDB.PhoneNumber,
	}
}

func (h handler) CreateCustomer(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        panic(err)
    }

	var customer models.Customer
    err = json.Unmarshal(body, &customer)
    if err != nil {
        panic(err)
    }
	err = customer.Validate()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error validating TCKN: %s", err),
			http.StatusBadRequest,
		)
		return
	}
	h.DB.Create(&customer)
	responseCustomer := ResponseCustomer(customer)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseCustomer)
}

func (h handler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []models.Customer{}
	h.DB.Find(&customers)
	responseCustomers := []Customer{}
	var customer models.Customer
	for _, customer = range customers{
		responseCustomer := ResponseCustomer(customer)
		responseCustomers = append(responseCustomers, responseCustomer)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func (h handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

	customer := models.Customer{}
	h.DB.Find(&customer, id)

	if customer.Id != id {
		http.Error(
			w,
			fmt.Sprintln("Customer Does Not Exists"),
			http.StatusBadRequest,
		)
	}

	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(customer)
}
