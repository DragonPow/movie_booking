package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/DragonPow/movie_booking/gen/proto/auth/v1"
	"github.com/DragonPow/movie_booking/internal/auth/config"
	"github.com/DragonPow/movie_booking/internal/auth/repository"
	"github.com/DragonPow/movie_booking/internal/auth/server"
)

type authServer struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	config     *config.Config
	repo       *repository.PostgresRepository
}

func newAuthServer(cfg *config.Config) (*authServer, error) {
	// Initialize repository
	repo, err := repository.NewPostgresRepository(cfg.Database)
	if err != nil {
		return nil, err
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Initialize auth server
	authSrv := server.NewAuthServer(repo, cfg)
	pb.RegisterAuthServiceServer(grpcServer, authSrv)

	// Enable reflection for grpc-curl
	reflection.Register(grpcServer)

	return &authServer{
		grpcServer: grpcServer,
		config:     cfg,
		repo:       repo,
	}, nil
}

func (s *authServer) start() error {
	// Start gRPC server
	grpcLis, err := net.Listen("tcp", s.config.Server.GRPCPort)
	if err != nil {
		return err
	}

	log.Printf("gRPC server starting on port %s", s.config.Server.GRPCPort)
	go s.grpcServer.Serve(grpcLis)

	// Create HTTP gateway
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Use localhost since gateway is in the same process
	endpoint := "localhost" + s.config.Server.GRPCPort
	if err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Printf("Failed to register gateway: %v", err)
		return err
	}

	// Start HTTP server
	s.httpServer = &http.Server{
		Addr:    s.config.Server.HTTPPort,
		Handler: mux,
	}

	log.Printf("HTTP gateway starting on port %s", s.config.Server.HTTPPort)
	return s.httpServer.ListenAndServe()
}

func (s *authServer) stop() {
	if s.repo != nil {
		if err := s.repo.Close(); err != nil {
			log.Printf("Error closing repository: %v", err)
		}
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down HTTP server: %v", err)
		}
	}
}
