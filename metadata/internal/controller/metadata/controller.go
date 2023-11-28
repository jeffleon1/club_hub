package metadata

import (
	"context"
	"errors"
	"time"

	metadataHttp "github.com/jeffleon1/club_hub/metadata/internal/gateway/http"
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"github.com/jeffleon1/club_hub/metadata/pkg/models"
)

var ErrNotFound = errors.New("not Found")

type repository interface {
	GetBy(ctx context.Context, key string, value interface{}) (*[]entities.Metadata, error)
	Create(ctx context.Context, entity *entities.Metadata) error
}

type Controller struct {
	repo    repository
	gateway *metadataHttp.Gateway
	retries int
}

func New(repo repository, gateway *metadataHttp.Gateway, retries int) *Controller {
	return &Controller{repo, gateway, retries}
}

func (c *Controller) Create(ctx context.Context, host string, franchiseID, companyID uint32) error {
	res, err := c.getGatewayWithRetry(ctx, host)
	if err != nil {
		return err
	}
	entity := res.ToMetadata(franchiseID, companyID)
	if err = c.repo.Create(ctx, entity); err != nil {
		return err
	}

	return nil
}

func (c *Controller) Get(ctx context.Context, companyID uint) (*[]entities.Metadata, error) {
	res, err := c.repo.GetBy(ctx, "company_id", companyID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) GetBy(ctx context.Context, key string, value uint32) (*[]entities.Metadata, error) {
	res, err := c.repo.GetBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) getGatewayWithRetry(ctx context.Context, host string) (*models.WebDomainDetails, error) {
	var res *models.WebDomainDetails
	var err error

	for i := 0; i < c.retries; i++ {
		res, err = c.gateway.Get(ctx, host)
		if err == nil && res != nil {
			break
		}

		time.Sleep(time.Duration(i) * time.Second)
	}

	return res, err
}
