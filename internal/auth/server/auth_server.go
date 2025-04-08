package server

import (
	authpb "github.com/DragonPow/movie_booking/gen/proto/auth/v1"
	"github.com/DragonPow/movie_booking/internal/auth/config"
	"github.com/DragonPow/movie_booking/internal/auth/repository"
)

// AuthServer handles authentication operations
type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	repo   repository.Repository
	config *config.Config
}

// NewAuthServer creates a new authentication server
func NewAuthServer(repo repository.Repository, cfg *config.Config) *AuthServer {
	return &AuthServer{
		repo:   repo,
		config: cfg,
	}
}
