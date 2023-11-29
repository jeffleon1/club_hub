package franchise

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	grpcGateway "github.com/jeffleon1/club_hub/franchise/internal/gateway/grpc"
	"github.com/jeffleon1/club_hub/franchise/pkg/entities"
	metadataEntity "github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/sirupsen/logrus"
)

var ErrNotFound = errors.New("not Found")

type franchiseRepository interface {
	GetBy(ctx context.Context, key string, value interface{}) (*[]entities.Franchise, error)
	Create(ctx context.Context, entity *entities.Franchise) error
	Update(ctx context.Context, entity *entities.Franchise, id uint) error
}

type countryRepository interface {
	Create(ctx context.Context, entity *entities.Country) error
	GetBy(ctx context.Context, key string, value interface{}) (*[]entities.Country, error)
}

type Controller struct {
	franchiseRepo franchiseRepository
	countryRepo   countryRepository
	gateway       *grpcGateway.Gateway
	retries       int
}

func New(franchiseRepo franchiseRepository, countryRepo countryRepository, gateway *grpcGateway.Gateway, retries int) *Controller {
	return &Controller{
		franchiseRepo,
		countryRepo,
		gateway,
		retries,
	}
}

func (c *Controller) Create(ctx context.Context, host string) error {
	host = strings.ToLower(host)
	franchise, err := c.franchiseRepo.GetBy(ctx, "url", host)
	if err != nil {
		return err
	}
	if len(*franchise) > 0 {
		logrus.Errorf("Error franchise already exists %s", host)
		return fmt.Errorf("Error franchise already exists %s", host)
	}
	res, err := c.getDomainInfo(host)
	if err != nil {
		logrus.Errorf("Error getting domain info %s", host)
		return err
	}
	countryEntity := entities.Country{
		Name: res.Registrant.Country,
	}
	countries, err := c.countryRepo.GetBy(ctx, "name", countryEntity.Name)
	if err != nil {
		return err
	}
	if len(*countries) == 0 {
		if err := c.countryRepo.Create(ctx, &countryEntity); err != nil {
			return err
		}
	} else {
		countryEntity = (*countries)[0]
	}
	entity := entities.WhoisToFranchise(res)
	entity.Location.CountryID = countryEntity.ID
	if err := c.franchiseRepo.Create(ctx, entity); err != nil {
		return err
	}
	if err := c.gateway.Create(ctx, host, uint32(entity.ID), uint32(entity.CompanyID)); err != nil {
		return err
	}
	return nil
}

func (c *Controller) Update(ctx context.Context, entity *entities.Franchise, id uint) error {
	return c.franchiseRepo.Update(ctx, entity, id)
}

func (c *Controller) Get(ctx context.Context, companyID uint) (*entities.Response, error) {
	franchises, err := c.franchiseRepo.GetBy(ctx, "company_id", companyID)
	if err != nil {
		return nil, err
	}
	metadatas, err := c.gateway.Get(ctx, uint32(companyID))
	if err != nil {
		return nil, err
	}

	return &entities.Response{
		Franchise: franchises,
		Metadata:  metadatas,
	}, nil
}

func (c *Controller) GetBy(ctx context.Context, key string, value string) (*entities.Response, error) {
	franchises, err := c.franchiseRepo.GetBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	var franchiseMetadata []*metadataEntity.Metadata
	for _, franchise := range *franchises {
		metadatas, err := c.gateway.GetBy(ctx, "franchise_id", uint32(franchise.ID))
		if err != nil {
			return nil, err
		}
		franchiseMetadata = append(franchiseMetadata, metadatas...)
	}

	if err != nil {
		return nil, err
	}

	return &entities.Response{
		Franchise: franchises,
		Metadata:  franchiseMetadata,
	}, nil
}

func (c *Controller) getDomainInfo(host string) (whoisparser.WhoisInfo, error) {
	var result whoisparser.WhoisInfo
	var err error

	for i := 0; i < c.retries; i++ {
		whoisRaw, err := whois.Whois(host)
		if err == nil {
			result, err = whoisparser.Parse(whoisRaw)
			if err == nil {
				break
			}
		}

		time.Sleep(time.Duration(i) * time.Second)
	}

	return result, err
}
