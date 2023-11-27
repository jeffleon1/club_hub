package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jeffleon1/club_hub/metadata/internal/gateway"
	"github.com/jeffleon1/club_hub/metadata/pkg/models"
)

type Gateway struct {
	BaseURL string
}

func New(BaseURL string) *Gateway {
	return &Gateway{BaseURL}
}

func (g *Gateway) Get(ctx context.Context, host string) (*models.WebDomainDetails, error) {
	url := g.BaseURL
	log.Printf("Calling service. Request: GET %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("host", host)
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response %v", resp)
	}
	var v *models.WebDomainDetails
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
