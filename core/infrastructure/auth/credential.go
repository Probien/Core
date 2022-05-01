package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthCustomClaims struct {
	Name      string `json:"name"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt time.Time
	jwt.StandardClaims
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
