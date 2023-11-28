package grpc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jeffleon1/club_hub/franchise/internal/controller/franchise"
	entityFranchise "github.com/jeffleon1/club_hub/franchise/pkg/entities"
	"github.com/jeffleon1/club_hub/gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedFranchiseServiceServer
	ctrl *franchise.Controller
}

func New(ctrl *franchise.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetFranchise(ctx context.Context, req *gen.GetFranchiseRequest) (*gen.GetFranchiseResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, uint(req.CompanyId))
	if err != nil && errors.Is(err, franchise.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var genFranchise []*gen.Franchise

	for _, franchise := range *m.Franchise {
		for _, metadata := range m.Metadata {
			if franchise.ID == metadata.FranchiseID {
				genFranchise = append(genFranchise, entityFranchise.FranchiseToProto(&franchise, metadata))
			}
		}
	}

	return &gen.GetFranchiseResponse{Franchise: genFranchise}, nil
}

func (h *Handler) GetFranchiseByFilter(ctx context.Context, req *gen.GetFranchiseByFilterRequest) (*gen.GetFranchiseResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	key := strings.ToLower(req.Key)
	if key != "name" && key != "url" && key != "email" {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Error filter key %s not allowed", key))
	}
	m, err := h.ctrl.GetBy(ctx, req.Key, req.Value)
	if err != nil && errors.Is(err, franchise.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var genFranchise []*gen.Franchise

	for _, franchise := range *m.Franchise {
		for _, metadata := range m.Metadata {
			if franchise.ID == metadata.FranchiseID {
				genFranchise = append(genFranchise, entityFranchise.FranchiseToProto(&franchise, metadata))
			}
		}
	}

	return &gen.GetFranchiseResponse{Franchise: genFranchise}, nil
}

func (h *Handler) CreateFranchise(ctx context.Context, req *gen.CreateFranchiseRequest) (*gen.CreateFranchiseResponse, error) {
	if req == nil || req.Host == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	if err := h.ctrl.Create(ctx, req.Host); err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	return &gen.CreateFranchiseResponse{Error: ""}, nil
}

func (h *Handler) UpdateFranchise(ctx context.Context, req *gen.UpdateFranchiseRequest) (*gen.UpdateFranchiseResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	logrus.Infof("Start Update")
	entity := entityFranchise.FranchiseFromProtoUpdate(req.Franchise)
	logrus.Infof("Continue")
	if err := h.ctrl.Update(ctx, entity, uint(req.Id)); err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	return &gen.UpdateFranchiseResponse{Error: ""}, nil
}
