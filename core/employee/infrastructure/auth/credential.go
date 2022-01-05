package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthCustomClaims struct {
	Name       string `json:"name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	IsAdmin    bool   `json:"is_admin"`
	CreatedAt  time.Time
	jwt.StandardClaims
}

func EncryptPassword(data []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return hash
}
