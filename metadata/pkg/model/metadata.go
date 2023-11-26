package model

type ProtocolType string

type Metadata struct {
	Protocol  ProtocolType `json:"protocol"`
	Endpoints []Endpoints  `json:"endpoints"`
}

type Endpoints struct {
	ServerName string `json:"server_name"`
	IpAdress   string `json:"ip_address"`
}
