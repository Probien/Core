package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(employee *domain.Employee, tokenizer chan<- string) {

	claims := &AuthCustomClaims{
		Name:      employee.Profile.Name,
		IsAdmin:   employee.IsAdmin,
		CreatedAt: time.Now(),
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

func GenerateSessionID(employee *domain.Employee, session chan<- SessionCredentials) {
	sessionID := uuid.NewV4()
	sessionClaims := SessionCredentials{ID: sessionID.String(), Username: employee.Email, ExpiresAt: time.Now().Add(time.Minute * 30)}
	sessionbBytes, _ := json.Marshal(sessionClaims)
	go config.Client.Set(context.Background(), sessionID.String(), string(sessionbBytes[:]), time.Minute*30)
	session <- sessionClaims
}

func ClearSessionID(c *gin.Context) {
	coockie, err := c.Cookie("SID")

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	}
	config.Client.Del(context.Background(), coockie)

}

func validateAndParseToken(encodedToken string, authCustomClaims *AuthCustomClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, authCustomClaims, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

}

func existCoockie(coockie string, checker chan<- bool) {
	var sessionID = SessionCredentials{}
	val := config.Client.Get(context.Background(), coockie).Val()
	json.Unmarshal([]byte(val), &sessionID)
	checker <- val != "" && coockie == sessionID.ID

}

func EncryptPassword(data []byte, ch chan<- []byte) {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	ch <- hash
}
