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

type Location struct {
    Id int `json:"id"`
    Name string `json:"name"`
    IsActive bool `json:"is_active"`
}

func ResponseLocation(location models.Location) Location {
	return Location{
		Id: location.Id,
		Name: location.Name,
		IsActive: location.IsActive,
	}
}

func (h handler) CreateLocation(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        panic(err)
    }

	var location models.Location
    err = json.Unmarshal(body, &location)
    if err != nil {
        panic(err)
    }
	h.DB.Create(&location)
	responseLocation := ResponseLocation(location)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseLocation)
}

func (h handler) GetAllLocations(w http.ResponseWriter, r *http.Request) {
	locations := []models.Location{}
	h.DB.Find(&locations)
	responseLocations := []Location{}
	var location models.Location
	for _, location = range locations{
		responseLocation := ResponseLocation(location)
		responseLocations = append(responseLocations, responseLocation)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)

}

func (h handler) GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

	location := models.Location{}
	h.DB.Find(&location, id)

	if location.Id != id {
		http.Error(
			w,
			fmt.Sprintln("Location Does Not Exists"),
			http.StatusBadRequest,
		)
	}

	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(location)
}
