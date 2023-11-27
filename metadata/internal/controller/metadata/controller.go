package metadata

import (
	"context"
	"errors"
	"fmt"

	metadataHttp "github.com/jeffleon1/club_hub/metadata/internal/gateway/http"
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
)

var ErrNotFound = errors.New("not Found")

type repository interface {
	GetBy(ctx context.Context, key string, value interface{}) (*[]entities.Metadata, error)
	Create(ctx context.Context, entity *entities.Metadata) error
}

type Controller struct {
	repo    repository
	gateway *metadataHttp.Gateway
}

func New(repo repository, gateway *metadataHttp.Gateway) *Controller {
	return &Controller{repo, gateway}
}

func (c *Controller) Create(ctx context.Context, host string) error {
	res, err := c.gateway.Get(ctx, host)
	if err != nil {
		return err
	}
	fmt.Println(res)
	entity := res.ToMetadata()
	fmt.Println("entity", entity)
	if err = c.repo.Create(ctx, entity); err != nil {
		return err
	}

	return nil
}

func (c *Controller) Get(ctx context.Context, franchiseID uint) (*[]entities.Metadata, error) {
	res, err := c.repo.GetBy(ctx, "franchise_id", franchiseID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
