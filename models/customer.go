package models

import (
	"github.com/go-playground/validator/v10"
)


type Customer struct {
    Id int `json:"id" gorm:"primaryKey"` 
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    TCKN string `json:"tckn" validate:"required,tckn"`
    BirthDate string `json:"birthDate"`
    PhoneNumber string `json:"phoneNumber"`
}

func TCKNValidator(fl validator.FieldLevel) bool {
    return fl.Field().Len() == 11 
}

func (customer *Customer) Validate() error {
    validate :=  validator.New()
    validate.RegisterValidation("tckn", TCKNValidator)
    return validate.Struct(customer)
}
