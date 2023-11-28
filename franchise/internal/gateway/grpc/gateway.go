package grpcGateway

import (
	"context"
	"fmt"

	"github.com/jeffleon1/club_hub/gen"
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	host string
	port string
}

// New creates a new GRPC gateway for a franchise metadata service
func New(host, port string) *Gateway {
	return &Gateway{host, port}
}

func (g *Gateway) Get(ctx context.Context, companyID uint32) ([]*entities.Metadata, error) {
	conn, err := ServiceConnection(ctx, g.host, g.port)
	if err != nil {
		logrus.Errorf("Error establishing connection with host=%s:%s error=%s", g.host, g.port, err.Error())
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	logrus.Infof("Sending GetMetadata company_id=%d", companyID)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{CompanyId: companyID})
	if err != nil {
		logrus.Errorf("Error GetMetadata company_id=%d error=%s", companyID, err.Error())
		return nil, err
	}
	var entityMetadata []*entities.Metadata

	for _, metadata := range resp.Metadata {
		entityMetadata = append(entityMetadata, entities.MetadataFromProto(metadata))
	}
	return entityMetadata, nil
}

func (g *Gateway) GetBy(ctx context.Context, key string, value uint32) ([]*entities.Metadata, error) {
	conn, err := ServiceConnection(ctx, g.host, g.port)
	if err != nil {
		logrus.Errorf("Error establishing connection with host=%s:%s error=%s", g.host, g.port, err.Error())
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	logrus.Infof("Sending GetMetadataByFilter key=%s value=%s", key, value)
	if err != nil {
		logrus.Errorf("Error GetMetadata key=%s value=%s %s", key, value, err.Error())
		return nil, err
	}
	resp, err := client.GetMetadataByFilter(ctx, &gen.GetMetadataByFilterRequest{Key: key, Value: value})
	if err != nil {
		logrus.Errorf("Error GetMetadata key=%s value=%s %s", key, value, err.Error())
		return nil, err
	}
	var entityMetadata []*entities.Metadata

	for _, metadata := range resp.Metadata {
		entityMetadata = append(entityMetadata, entities.MetadataFromProto(metadata))
	}
	return entityMetadata, nil
}

func (g *Gateway) Create(ctx context.Context, host string, franchiseID, companyID uint32) error {
	conn, err := ServiceConnection(ctx, g.host, g.port)
	if err != nil {
		logrus.Errorf("Error establishing connection with host=%s:%s error=%s", g.host, g.port, err.Error())
		return err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	logrus.Infof("Sending CreateMetadata host=%s franchise_id=%d company_id=%d", host, franchiseID, companyID)
	_, err = client.CreateMetadata(ctx, &gen.CreateMetadataRequest{Host: host, FranchiseId: franchiseID, CompanyId: companyID})
	if err != nil {
		logrus.Errorf("Error CreateMetadata host=%s franchise_id=%d company_id=%d error=%s", host, franchiseID, companyID, err.Error())
		return err
	}
	return nil
}

func ServiceConnection(ctx context.Context, host, port string) (*grpc.ClientConn, error) {
	return grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
