package models


type Location struct {
    Id int `json:"id" gorm:"primary_key;auto_increment;not_null"` 
    Name string `json:"name"`
    IsActive bool `json:"isActive"`
}

// func isActive(fl validator.FieldLevel) bool {
// 	return fl.Field() == true
// }

// func (location *Location) Validate() error {
//     validate :=  validator.New()
//     validate.RegisterValidation("isActive", isActive)
//     return validate.Struct(location)
// }
