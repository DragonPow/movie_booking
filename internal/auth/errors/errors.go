package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrInvalidCredentials is returned when email/password combination is invalid
	ErrInvalidCredentials = status.Error(codes.Unauthenticated, "invalid credentials")

	// ErrInvalidToken is returned when JWT token is invalid
	ErrInvalidToken = status.Error(codes.Unauthenticated, "invalid token")

	// ErrMissingField is returned when a required field is missing
	ErrMissingField = status.Error(codes.InvalidArgument, "missing required field")

	// ErrInvalidEmail is returned when email format is invalid
	ErrInvalidEmail = status.Error(codes.InvalidArgument, "invalid email format")

	// ErrInternalServer is returned when an internal error occurs
	ErrInternalServer = status.Error(codes.Internal, "internal server error")

	// ErrUserExists is returned when attempting to register with an existing email
	ErrUserExists = status.Error(codes.AlreadyExists, "user already exists")

	// ErrInvalidUsername is returned when username format is invalid
	ErrInvalidUsername = status.Error(codes.InvalidArgument, "invalid username format")

	// ErrInvalidPassword is returned when password does not meet requirements
	ErrInvalidPassword = status.Error(codes.InvalidArgument, "invalid password format")
)
