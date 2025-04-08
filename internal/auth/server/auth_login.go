package server

import (
	"context"

	authpb "github.com/DragonPow/movie_booking/gen/proto/auth/v1"
	"github.com/DragonPow/movie_booking/internal/auth/errors"
	"github.com/DragonPow/movie_booking/internal/auth/repository"
	"github.com/DragonPow/movie_booking/internal/auth/validation"
	"golang.org/x/crypto/bcrypt"
)

// Login authenticates a user and returns a JWT token
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// Validate inputs
	if err := validation.ValidateEmail(req.Email); err != nil {
		return nil, err
	}
	if req.Password == "" {
		return nil, errors.ErrMissingField
	}

	// Get user by email
	normalizedEmail := validation.NormalizeEmail(req.Email)
	user, err := s.repo.GetUserByEmail(ctx, normalizedEmail)
	if err != nil {
		if repository.IsNotFoundError(err) {
			return nil, errors.ErrInvalidCredentials
		}
		return nil, errors.ErrInternalServer
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	// Generate token
	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return &authpb.LoginResponse{
		UserId: user.ID.String(),
		Token:  token,
	}, nil
}
