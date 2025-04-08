package middleware

import (
	"context"
	"strings"

	"github.com/DragonPow/movie_booking/internal/auth/server"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor provides authentication for gRPC services
type AuthInterceptor struct {
	authServer *server.AuthServer
}

// NewAuthInterceptor creates a new auth interceptor
func NewAuthInterceptor(authServer *server.AuthServer) *AuthInterceptor {
	return &AuthInterceptor{
		authServer: authServer,
	}
}

// Unary returns a unary server interceptor for authentication
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Skip auth for login and register
		if info.FullMethod == "/auth.AuthService/Login" || info.FullMethod == "/auth.AuthService/Register" {
			return handler(ctx, req)
		}

		userID, err := i.authorize(ctx)
		if err != nil {
			return nil, err
		}

		// Add user ID to context
		newCtx := context.WithValue(ctx, "user_id", userID)
		return handler(newCtx, req)
	}
}

// Stream returns a stream server interceptor for authentication
func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		userID, err := i.authorize(ss.Context())
		if err != nil {
			return err
		}

		// Wrap stream with new context containing user ID
		wrapper := &authStreamWrapper{
			ServerStream: ss,
			ctx:          context.WithValue(ss.Context(), "user_id", userID),
		}
		return handler(srv, wrapper)
	}
}

func (i *AuthInterceptor) authorize(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return uuid.Nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.TrimPrefix(values[0], "Bearer ")
	userID, err := i.authServer.ValidateTokenString(token)
	if err != nil {
		return uuid.Nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}

	return userID, nil
}

// authStreamWrapper wraps grpc.ServerStream to modify its context
type authStreamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *authStreamWrapper) Context() context.Context {
	return w.ctx
}
