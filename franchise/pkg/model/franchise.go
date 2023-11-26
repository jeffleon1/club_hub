package model

import (
	"time"

	"gorm.io/gorm"
)

type Franchise struct {
	gorm.Model
	Url         string    `json:"url"`
	CreateDate  time.Time `json:"create_date"`
	ExpiresDate time.Time `json:"expires_date"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Location    Location  `json:"location"`
}

type Location struct {
	gorm.Model
	City    City   `json:"city"`
	Address string `json:"address"`
}

type Country struct {
	gorm.Model
	Name string `json:"name"`
}

type City struct {
	gorm.Model
	Name    string  `json:"name"`
	Country Country `json:"country"`
}
