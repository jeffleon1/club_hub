package entities

import (
	"context"

	"github.com/jeffleon1/club_hub/gen"
	"gorm.io/gorm"
)

type ProtocolType string

type Metadata struct {
	gorm.Model
	Protocol     ProtocolType `json:"protocol"`
	Availability string       `json:"availability"`
	Endpoints    []Endpoint   `json:"endpoints" gorm:"foreignKey:MetadataID"`
	FranchiseID  uint         `json:"franchise_id"`
}

type Endpoint struct {
	gorm.Model
	IPAddress  string `json:"ipAddress"`
	ServerName string `json:"serverName"`
	MetadataID uint   `json:"-" gorm:"index"`
}

func MetadataToProto(m *Metadata) *gen.Metadata {
	var endpoints []*gen.Endpoint

	for _, endpoint := range m.Endpoints {
		endpoints = append(endpoints, &gen.Endpoint{
			ServerName: endpoint.ServerName,
			IpAddress:  endpoint.IPAddress,
		})
	}

	return &gen.Metadata{
		Protocol:     string(m.Protocol),
		Availability: m.Availability,
		Endpoints:    endpoints,
		FranchiseId:  uint32(m.FranchiseID),
	}
}

// MetadataFromProto converts a generated proto counterpart
// into a Metadata struct.
func MetadataFromProto(m *gen.Metadata) *Metadata {
	var endpoints []Endpoint

	for _, endpoint := range m.Endpoints {
		endpoints = append(endpoints, Endpoint{
			ServerName: endpoint.ServerName,
			IPAddress:  endpoint.IpAddress,
		})
	}

	return &Metadata{
		Protocol:     ProtocolType(m.Protocol),
		Availability: m.Availability,
		Endpoints:    endpoints,
		FranchiseID:  uint(m.FranchiseId),
	}
}

type Repository[T interface{}] interface {
	Create(ctx context.Context, entity *T) error
	GetAll(ctx context.Context) (*[]T, error)
	GetByID(ctx context.Context, id uint) (*T, error)
	Update(ctx context.Context, entity *T, id uint) error
	Delete(ctx context.Context, id uint) error
	GetBy(ctx context.Context, key string, value interface{}) (*[]T, error)
}
