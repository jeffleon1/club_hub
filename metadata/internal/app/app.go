package app

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jeffleon1/club_hub/gen"
	pb "github.com/jeffleon1/club_hub/gen"
	"github.com/jeffleon1/club_hub/metadata/config"
	"github.com/jeffleon1/club_hub/metadata/config/db"
	"github.com/jeffleon1/club_hub/metadata/internal/controller/metadata"
	metadataGateway "github.com/jeffleon1/club_hub/metadata/internal/gateway/http"
	grpcHandler "github.com/jeffleon1/club_hub/metadata/internal/handler/grpc"
	"github.com/jeffleon1/club_hub/metadata/internal/repository/posgrest"
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
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
	repository := posgrest.New[entities.Metadata](a.db, a.config.DB.RELATIONS)
	gateway := metadataGateway.New(a.config.GATEWAY.URL)
	ctrl := metadata.New(repository, gateway, a.config.APP.RETRIES)
	handler := grpcHandler.New(ctrl)
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	logrus.Info("Metadata GRPC service working in port 8081")
	gen.RegisterMetadataServiceServer(srv, handler)
	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	err = pb.RegisterMetadataServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: mux,
	}

	log.Printf("Serving gRPC-Gateway on connection in 8091")
	log.Fatalln(gwServer.ListenAndServe())
}
