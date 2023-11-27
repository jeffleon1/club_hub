package models

type ProtocolType string

type Metadata struct {
	Protocol     ProtocolType `json:"protocol"`
	Availability string       `json:"availability"`
	Endpoints    Endpoint     `json:"endpoints"`
	FranchiseID  uint         `json:"franchise_id"`
}
