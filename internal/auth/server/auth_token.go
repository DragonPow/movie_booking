package server

import (
	"context"
	"strings"
	"time"

	authpb "github.com/DragonPow/movie_booking/gen/proto/auth/v1"
	"github.com/DragonPow/movie_booking/internal/auth/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func (s *AuthServer) GenerateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(s.config.JWT.Expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", errors.ErrInternalServer
	}

	return signedToken, nil
}

func (s *AuthServer) ValidateTokenString(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidToken
		}
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil {
		return uuid.Nil, errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return uuid.Nil, errors.ErrInvalidToken
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, errors.ErrInvalidToken
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.ErrInvalidToken
	}

	return userID, nil
}

// ValidateTokenFromContext extracts and validates token from context
func (s *AuthServer) ValidateTokenFromGRPC(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, errors.ErrInvalidToken
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return uuid.Nil, errors.ErrInvalidToken
	}

	bearerToken := authHeader[0]
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return uuid.Nil, errors.ErrInvalidToken
	}

	token := strings.TrimPrefix(bearerToken, "Bearer ")
	return s.ValidateTokenString(token)
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	userID, err := s.ValidateTokenFromGRPC(ctx)
	if err != nil {
		return nil, err
	}

	// Get user details from repository
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &authpb.ValidateTokenResponse{
		UserId: userID.String(),
		Email:  user.Email,
	}, nil
}
