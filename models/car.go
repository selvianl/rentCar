package models

type Car struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	FuelType string `json:"fuelType"`
	GearType string `json:"gearType"`
	VendorID int `json:"vendorId"`
	Vendor Vendor `gorm:"foreignKey:VendorRefer"` 
}