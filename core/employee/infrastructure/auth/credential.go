package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

/*
type Credentials struct {
	secretKey string
	issue     string
}

func GetJWTCredentials() *Credentials {
	return &Credentials{
		secretKey: base64.StdEncoding.EncodeToString([]byte("EQVJ7UM8xJNcfsaxs$aw3Es2Z@8ewegzxZ531C$^bhEoMq!%fe")),
		issue:     "Probien",
	}
}
*/
