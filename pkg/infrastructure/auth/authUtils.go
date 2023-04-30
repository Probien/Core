package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(employee *domain.Employee, tokenizer chan<- string) {

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
}

func GenerateSessionID(employee *domain.Employee, session chan<- SessionCredential) {
	sessionID := uuid.NewV4()
	sessionClaims := SessionCredential{
		ID:        sessionID.String(),
		Username:  employee.Email,
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}
	sessionBytes, err := json.Marshal(sessionClaims)
	if err != nil {
		fmt.Errorf("error marshalling session: %v", err)
	}
	cmd := config.Client.Set(context.Background(), sessionID.String(), string(sessionBytes[:]), time.Minute*30)
	if err := cmd.Err(); err != nil {
		fmt.Errorf("error writing session to Redis: %v", err)
	}
	session <- sessionClaims
}

func ClearSessionID(c *gin.Context) error {
	cookie, err := c.Cookie("SID")

	if err != nil {
		return err
	}

	config.Client.Del(context.Background(), cookie)
	return nil
}

func validateAndParseToken(encodedToken string, authCustomClaims *CustomClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, authCustomClaims, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

}

func existCookie(cookie string, checker chan<- bool) {
	var sessionID = SessionCredential{}
	val := config.Client.Get(context.Background(), cookie).Val()
	err := json.Unmarshal([]byte(val), &sessionID)
	if err != nil {
		fmt.Errorf("error getting session from Redis: %v", err)
	}
	checker <- val != "" && cookie == sessionID.ID

}

func EncryptPassword(data []byte, ch chan<- []byte) {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		fmt.Errorf("error encrypting password: %v", err)
	}
	ch <- hash
}

func checkAuthorities(authorities []string, authCustomClaims *CustomClaims) bool {
	for i := 0; i < len(authorities); i++ {
		for _, role := range authCustomClaims.Roles {
			if role == authorities[i] {
				return true
			}
		}
	}
	return false
}
