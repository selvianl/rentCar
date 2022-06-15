package models

type Office struct {
	Id int `json:"id" gorm:"primaryKey"`
	LocationRefer int `json:"location_id"`
	Location Location `gorm:"foreignKey:LocationRefer"`
	VendorRefer int `json:"vendor_id"`
	Vendor Vendor `gorm:"foreignKey:VendorRefer"`
}