package component

import (
	"fmt"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator struct {
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (a *Authenticator) GenerateToken(employee *dto.Employee, tokenizer chan<- string) {

	roles := make(map[string]string)

	for k, v := range employee.Roles {
		roles["role_"+strconv.Itoa(k)] = v.Role.RoleName
	}

	claims := &CustomClaims{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			Issuer:    "Probien",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(int(employee.ID)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
	if err != nil {
		panic(err)
	}
	tokenizer <- token
	close(tokenizer)
	return
}

func (a *Authenticator) ValidateAndParseToken(encodedToken string, authCustomClaims *CustomClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, authCustomClaims, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

}

func (a *Authenticator) EncryptPassword(data []byte, ch chan<- []byte) {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		fmt.Errorf("error encrypting password: %v", err)
	}
	ch <- hash
	close(ch)
	return
}

func (a *Authenticator) CheckAuthorities(authorities []string, authCustomClaims *CustomClaims) bool {
	for i := 0; i < len(authorities); i++ {
		for _, role := range authCustomClaims.Roles {
			if role == authorities[i] {
				return true
			}
		}
	}
	return false
}
