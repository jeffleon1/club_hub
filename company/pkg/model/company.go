package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Owner       Owner       `json:"owner"`
	Information Information `json:"information"`
}

type Owner struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Contact   Contact `json:"contact"`
}

type Contact struct {
	gorm.Model
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Information struct {
	gorm.Model
	Name      string `json:"name"`
	TaxNumber string `json:"tax_number"`
}
