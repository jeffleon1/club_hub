package app

import (
	"fmt"
	"log"
	"net"

	"github.com/jeffleon1/club_hub/franchise/config"
	"github.com/jeffleon1/club_hub/franchise/config/db"
	"github.com/jeffleon1/club_hub/franchise/internal/controller/franchise"
	grpcGateway "github.com/jeffleon1/club_hub/franchise/internal/gateway/grpc"
	grpcHandler "github.com/jeffleon1/club_hub/franchise/internal/handler/grpc"
	"github.com/jeffleon1/club_hub/franchise/internal/repository/posgrest"
	"github.com/jeffleon1/club_hub/franchise/pkg/entities"
	"github.com/jeffleon1/club_hub/gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	config *config.Config
}

func (a *App) Initialize(cfg *config.Config) {
	a.config = cfg
	a.db = db.Connect(cfg)
	if err := db.MigrateDB(a.db); err != nil {
		panic(err)
	}
	franchiseRepository := posgrest.New[entities.Franchise](a.db, "Locations")
	countryRepository := posgrest.New[entities.Country](a.db, "")
	gateway := grpcGateway.New(a.config.GATEWAY.HOST, a.config.GATEWAY.PORT)
	ctrl := franchise.New(franchiseRepository, countryRepository, gateway, a.config.APP.RETRIES)
	handler := grpcHandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", a.config.APP.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	logrus.Infof("Franchise GRPC service working in port %s", a.config.APP.PORT)
	gen.RegisterFranchiseServiceServer(srv, handler)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
