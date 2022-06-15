package models

import (
	"github.com/lib/pq"
)

type Vendor struct {
	Id int `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Name string `json:"name"`
	WorkingHours string`json:"workingHours"`
	ClosingHours string `json:"closingHours"`
	WorkingDays pq.StringArray `gorm:"type:text[]"`
	LocationRefer int `json:"location_id"`
	Location Location `gorm:"foreignKey:LocationRefer"`
}
