package grpc

import (
	"context"
	"errors"

	"github.com/jeffleon1/club_hub/gen"
	"github.com/jeffleon1/club_hub/metadata/internal/controller/metadata"
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedMetadataServiceServer
	ctrl *metadata.Controller
}

func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, uint(req.CompanyId))
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var genMetadata []*gen.Metadata

	for _, metadata := range *m {
		genMetadata = append(genMetadata, entities.MetadataToProto(&metadata))
	}

	return &gen.GetMetadataResponse{Metadata: genMetadata}, nil
}

func (h *Handler) GetMetadataByFilter(ctx context.Context, req *gen.GetMetadataByFilterRequest) (*gen.GetMetadataResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.GetBy(ctx, req.Key, req.Value)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var genMetadata []*gen.Metadata

	for _, metadata := range *m {
		genMetadata = append(genMetadata, entities.MetadataToProto(&metadata))
	}

	return &gen.GetMetadataResponse{Metadata: genMetadata}, nil
}

func (h *Handler) CreateMetadata(ctx context.Context, req *gen.CreateMetadataRequest) (*gen.CreateMetadataResponse, error) {
	if req == nil || req.Host == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	if err := h.ctrl.Create(ctx, req.Host, req.FranchiseId, req.CompanyId); err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	return &gen.CreateMetadataResponse{Error: ""}, nil
}
