package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jeffleon1/club_hub/franchise/config"
	"github.com/jeffleon1/club_hub/franchise/config/db"
	"github.com/jeffleon1/club_hub/franchise/internal/controller/franchise"
	grpcGateway "github.com/jeffleon1/club_hub/franchise/internal/gateway/grpc"
	grpcHandler "github.com/jeffleon1/club_hub/franchise/internal/handler/grpc"
	"github.com/jeffleon1/club_hub/franchise/internal/repository/posgrest"
	"github.com/jeffleon1/club_hub/franchise/pkg/entities"
	"github.com/jeffleon1/club_hub/gen"
	pb "github.com/jeffleon1/club_hub/gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	err = pb.RegisterFranchiseServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Printf("Serving gRPC-Gateway on connection in 8090")
	log.Fatalln(gwServer.ListenAndServe())
}
