package model

import "time"

type DomainDetails struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CreateDate  time.Time `json:"create_date"`
	ExpiresDate time.Time `json:"expires_date"`
}
