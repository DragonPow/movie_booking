package validation

import (
	"net/mail"
	"strings"
	"unicode"

	"github.com/DragonPow/movie_booking/internal/auth/errors"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 72 // bcrypt max length
)

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) error {
	if email == "" {
		return errors.ErrMissingField
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.ErrInvalidEmail
	}
	return nil
}

// ValidatePassword checks if the password meets security requirements
func ValidatePassword(password string) error {
	if password == "" {
		return errors.ErrMissingField
	}
	if len(password) < minPasswordLength {
		return errors.ErrInvalidPassword
	}
	if len(password) > maxPasswordLength {
		return errors.ErrInvalidPassword
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !(hasUpper && hasLower && hasNumber && hasSpecial) {
		return errors.ErrInvalidPassword
	}

	return nil
}

// ValidateUsername checks if the username is valid
func ValidateUsername(username string) error {
	if username == "" {
		return errors.ErrMissingField
	}
	if len(username) < 3 || len(username) > 50 {
		return errors.ErrInvalidUsername
	}
	return nil
}

// NormalizeEmail normalizes email for consistent comparison
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
