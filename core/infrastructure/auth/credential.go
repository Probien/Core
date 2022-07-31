package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthCustomClaims struct {
	Name    string            `json:"name"`
	IsAdmin bool              `json:"is_admin"`
	Roles   map[string]string `json:"roles"`
	jwt.RegisteredClaims
}

type SessionCredential struct {
	ID        string
	Username  string
	ExpiresAt time.Time
}

type LoginCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
