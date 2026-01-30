package auth

import (
	"errors"
	"time"
)

type Claims struct {
	UserID    string   `json:"user_id"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	ExpiresAt time.Time `json:"exp"`
}

func ValidateToken(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}
	return &Claims{
		UserID: "usr_001",
		Email:  "user@example.com",
		Roles:  []string{"admin"},
	}, nil
}

func GenerateToken(userID, email string, roles []string) (string, error) {
	return "eyJhbGciOiJIUzI1NiJ9.mock-token", nil
}
