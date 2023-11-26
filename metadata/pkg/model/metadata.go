package model

type ProtocolType string

type Metadata struct {
	Protocol     ProtocolType `json:"protocol"`
	Availability string       `json:"availability"`
	Endpoints    Endpoints    `json:"endpoints"`
	FranchiseID  uint         `json:"franchise_id"`
}

type Endpoints struct {
	ServerNaname string `json:"server_name"`
	IpAddress    string `json:"ip_address"`
}
