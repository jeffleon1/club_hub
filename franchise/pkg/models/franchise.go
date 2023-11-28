package models

import (
	"time"

	entityFranchise "github.com/jeffleon1/club_hub/franchise/pkg/entities"
	entityMetadata "github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"github.com/jeffleon1/club_hub/metadata/pkg/models"
)

type Franchise struct {
	ID          uint            `json:"-"`
	URL         string          `json:"url"`
	CreateDate  time.Time       `json:"create_date"`
	ExpiresDate time.Time       `json:"expires_date"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	Location    Location        `json:"location" gorm:"foreignKey:LocationID"`
	CompanyID   uint            `json:"company_id"`
	Metadata    models.Metadata `json:"metadata"`
}

type Location struct {
	City       string  `json:"city"`
	Address    string  `json:"address"`
	LocationID uint    `json:"-" gorm:"index"`
	CountryID  uint    `json:"country_id"`
	Country    Country `json:"country" gorm:"foreignKey:CountryID"`
}

type Country struct {
	Name string `json:"name"`
}

func ToModelsFranchise(entity entityFranchise.Franchise, metadata entityMetadata.Metadata) *Franchise {
	var modelEndpoints []models.Endpoint
	for _, endpoint := range metadata.Endpoints {
		modelEndpoints = append(modelEndpoints, models.Endpoint{
			IPAddress:  endpoint.IPAddress,
			ServerName: endpoint.ServerName,
		})
	}

	return &Franchise{
		URL:         entity.URL,
		CreateDate:  entity.CreateDate,
		ExpiresDate: entity.ExpiresDate,
		Name:        entity.Name,
		Email:       entity.Email,
		CompanyID:   entity.CompanyID,
		Location: Location{
			City:    entity.Location.City,
			Address: entity.Location.Address,
			Country: Country{
				Name: entity.Location.Country.Name,
			},
		},
		Metadata: models.Metadata{
			Protocol:     models.ProtocolType(metadata.Protocol),
			Availability: metadata.Availability,
			FranchiseID:  metadata.FranchiseID,
			Endpoints:    modelEndpoints,
		},
	}
}
