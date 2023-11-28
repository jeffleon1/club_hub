package models

import (
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
)

type WebDomainDetails struct {
	Host            string     `json:"host"`
	Port            int64      `json:"port"`
	Protocol        string     `json:"protocol"`
	IsPublic        bool       `json:"isPublic"`
	Status          string     `json:"status"`
	StartTime       int64      `json:"startTime"`
	EngineVersion   string     `json:"engineVersion"`
	CriteriaVersion string     `json:"criteriaVersion"`
	Endpoints       []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	ServerName    string `json:"serverName"`
	StatusMessage string `json:"statusMessage"`
	Delegation    int64  `json:"delegation"`
}

func (w *WebDomainDetails) ToMetadata(franchiseID, companyID uint32) *entities.Metadata {
	var endpoints []entities.Endpoint
	for _, endpoint := range w.Endpoints {
		endpoints = append(endpoints, entities.Endpoint{
			IPAddress:  endpoint.IPAddress,
			ServerName: endpoint.ServerName,
		})
	}

	return &entities.Metadata{
		Protocol:     entities.ProtocolType(w.Protocol),
		Availability: w.Status,
		Endpoints:    endpoints,
		FranchiseID:  uint(franchiseID),
		CompanyID:    uint(companyID),
	}
}
