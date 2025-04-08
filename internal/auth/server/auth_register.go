package server

import (
	"context"

	authpb "github.com/DragonPow/movie_booking/gen/proto/auth/v1"
	"github.com/DragonPow/movie_booking/internal/auth/errors"
	"github.com/DragonPow/movie_booking/internal/auth/repository"
	"github.com/DragonPow/movie_booking/internal/auth/validation"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	// Validate inputs
	if err := validation.ValidateUsername(req.Username); err != nil {
		return nil, err
	}
	if err := validation.ValidateEmail(req.Email); err != nil {
		return nil, err
	}
	if err := validation.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// Normalize email
	normalizedEmail := validation.NormalizeEmail(req.Email)

	// Check if email already exists
	existingUser, err := s.repo.GetUserByEmail(ctx, normalizedEmail)
	if err != nil && !repository.IsNotFoundError(err) {
		return nil, errors.ErrInternalServer
	}
	if existingUser.ID != uuid.Nil {
		return nil, errors.ErrUserExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	// Create user
	userID := uuid.New()
	params := repository.CreateUserParams{
		ID:       userID,
		Username: req.Username,
		Email:    normalizedEmail,
		Password: string(hashedPassword),
	}
	if err := s.repo.CreateUser(ctx, params); err != nil {
		return nil, errors.ErrInternalServer
	}

	// Generate token
	token, err := s.GenerateToken(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return &authpb.RegisterResponse{
		UserId: userID.String(),
		Token:  token,
	}, nil
}
