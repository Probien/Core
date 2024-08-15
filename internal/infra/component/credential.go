package component

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Roles map[string]string `json:"roles"`
	jwt.RegisteredClaims
}

type SessionCredential struct {
	ID        string
	Username  string
	ExpiresAt time.Time
}

type Credential struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
