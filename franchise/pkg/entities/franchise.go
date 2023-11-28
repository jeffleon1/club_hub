package entities

import (
	"time"

	"github.com/jeffleon1/club_hub/gen"
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
	whoisparser "github.com/likexian/whois-parser"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Response struct {
	Franchise *[]Franchise
	Metadata  []*entities.Metadata
}

type Franchise struct {
	gorm.Model
	URL         string    `json:"url"`
	CreateDate  time.Time `json:"create_date"`
	ExpiresDate time.Time `json:"expires_date"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Location    Location  `json:"location" gorm:"foreignKey:LocationID"`
	CompanyID   uint      `json:"company_id"`
}

type Location struct {
	gorm.Model
	City       string  `json:"city"`
	Address    string  `json:"address"`
	LocationID uint    `json:"-" gorm:"index"`
	CountryID  uint    `json:"country_id"`
	Country    Country `json:"country" gorm:"foreignKey:CountryID"`
}

type Country struct {
	gorm.Model
	Name string `json:"name"`
}

func WhoisToFranchise(whoisInfo whoisparser.WhoisInfo) *Franchise {
	return &Franchise{
		URL:         whoisInfo.Domain.Domain,
		CreateDate:  *whoisInfo.Domain.CreatedDateInTime,
		ExpiresDate: *whoisInfo.Domain.ExpirationDateInTime,
		Name:        whoisInfo.Domain.Name,
		Email:       whoisInfo.Registrant.Email,
		CompanyID:   0,
		Location: Location{
			City:    whoisInfo.Registrant.City,
			Address: whoisInfo.Registrant.Street,
		},
	}
}

func FranchiseToProto(f *Franchise, m *entities.Metadata) *gen.Franchise {
	return &gen.Franchise{
		Url:         f.URL,
		CreateDate:  timestamppb.New(f.CreateDate),
		ExpiresDate: timestamppb.New(f.ExpiresDate),
		Name:        f.Name,
		Email:       f.Email,
		CompanyId:   uint32(f.CompanyID),
		Location: &gen.Location{
			City:    f.Location.City,
			Address: f.Location.Address,
			Country: &gen.Country{
				Name: f.Location.Country.Name,
			},
		},
		Metadata: entities.MetadataToProto(m),
	}
}

func FranchiseFromProto(m *gen.Franchise) *Franchise {
	return &Franchise{
		URL:         m.Url,
		CreateDate:  m.CreateDate.AsTime(),
		ExpiresDate: m.ExpiresDate.AsTime(),
		Name:        m.Name,
		Email:       m.Email,
		Location: Location{
			City:    m.Location.City,
			Address: m.Location.Address,
			Country: Country{
				Name: m.Location.Country.Name,
			},
		},
		CompanyID: uint(m.CompanyId),
	}
}

func FranchiseFromProtoUpdate(m *gen.Franchise) *Franchise {
	var existingFranchise Franchise
	if m.Url != "" {
		existingFranchise.URL = m.Url
	}
	if m.CreateDate != nil {
		existingFranchise.CreateDate = m.CreateDate.AsTime()
	}
	if m.ExpiresDate != nil {
		existingFranchise.ExpiresDate = m.ExpiresDate.AsTime()
	}
	if m.Name != "" {
		existingFranchise.Name = m.Name
	}
	if m.Email != "" {
		existingFranchise.Email = m.Email
	}
	if m.Location != nil {
		if m.Location.City != "" {
			existingFranchise.Location.City = m.Location.City
		}
		if m.Location.Address != "" {
			existingFranchise.Location.Address = m.Location.Address
		}
		if m.Location.Country != nil && m.Location.Country.Name != "" {
			existingFranchise.Location.Country.Name = m.Location.Country.Name
		}
	}
	if m.CompanyId != 0 {
		existingFranchise.CompanyID = uint(m.CompanyId)
	}

	return &existingFranchise
}
