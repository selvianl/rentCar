package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentCar/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"gorm.io/gorm"
)
type Vendor struct {
	Id int `json:"id"`
	Name string `json:"name"`
	WorkingHours string`json:"workingHours"`
	ClosingHours string `json:"closingHours"`
	WorkingDays pq.StringArray `json:"working_days"`
	Location Location  `json:"location"`
}
func ResponseVendor(vendorDB models.Vendor, location Location) Vendor{
	return Vendor{
		Id: vendorDB.Id,
		Name: vendorDB.Name,
		WorkingHours: vendorDB.WorkingHours,
		ClosingHours: vendorDB.ClosingHours,
		WorkingDays: vendorDB.WorkingDays,
		Location: location,
	}
}

func (h handler) GetAllVendors(w http.ResponseWriter, r *http.Request) {
	vendors := []models.Vendor{}
	h.DB.Find(&vendors)
	responseVendors := []Vendor{}
	location := models.Location{}
	for _, vendor := range vendors {
		h.DB.Find(&location, vendor.LocationRefer)
		responseVendor := ResponseVendor(vendor, ResponseLocation(location))
		responseVendors = append(responseVendors, responseVendor)
	}
	if result := h.DB.Find(&vendors); result.Error !=nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendors)
}

func (h handler) CreateVendor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
        panic(err)
    }

	var vendor models.Vendor
	err = json.Unmarshal(body, &vendor)
	if err != nil {
		panic(err)
	}

	var location models.Location
	if err := h.DB.Find(&location, vendor.LocationRefer).Error 
	err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		panic(err)
	}
	h.DB.Create(&vendor)
	ResponseLocation := ResponseLocation(location)
	responseVendor := ResponseVendor(vendor, ResponseLocation)

	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(responseVendor)

}

func (h handler) GetVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	vendor := models.Vendor{}
	h.DB.Find(&vendor, id)
	if vendor.Id != id {
		http.Error(
			w,
			fmt.Sprintln("Location Does Not Exists"),
			http.StatusBadRequest,
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(vendor)

}