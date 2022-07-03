package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	token, err := t.SignedString([]byte("DPzN3tMBaKsAPxvq8hWfaBHu5oeoj4bioNMQ6NzBSifkTthYAcoM67NzWTaZbPSDhGTkZhsdxyvYmNALanSoa3MH8CBW6Auv"))
	if err != nil {
		panic(err)
	}
	tokenizer <- token
}

func GenerateSessionID(employee *domain.Employee, session chan<- SessionCredentials) {
	sessionID := uuid.NewV4()
	sessionClaims := SessionCredentials{ID: sessionID.String(), Username: employee.Email, ExpiresAt: time.Now().Add(time.Minute * 1)}
	sessionbBytes, _ := json.Marshal(sessionClaims)
	go config.Client.Set(context.Background(), sessionID.String(), string(sessionbBytes[:]), time.Minute*1)
	session <- sessionClaims
}

func JwtAuth(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		checker := make(chan string, 1)
		coockie, _ := c.Cookie("SID")
		go existCoockie(coockie, checker)
		authHeader := c.GetHeader("Authorization")
		data := AuthCustomClaims{}

		if len(authHeader) > 0 && authHeader != "Bearer" {
			splitToken := strings.Split(authHeader, "Bearer")
			encodedToken := strings.TrimSpace(splitToken[1])
			token, _ := validateAndParseToken(encodedToken, &data)

			if token.Valid && data.IsAdmin == isAdmin && <-checker == coockie {

				//extract user_id from parsed token for stored procedures
				user_id, _ := strconv.Atoi(data.RegisteredClaims.Subject)
				//set user_id to request only for this context(request)
				c.Set("user_id", user_id)
				c.Next()

			} else {
				c.JSON(http.StatusUnauthorized, common.Response{Status: http.StatusUnauthorized, Message: "Authorization is required", Data: "Unauthorized, valid token is required"})
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.JSON(http.StatusUnauthorized, common.Response{Status: http.StatusUnauthorized, Message: "Authorization is required", Data: "Token is not present in the request header"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EncryptPassword(data []byte, ch chan<- []byte) {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	ch <- hash
}

func validateAndParseToken(encodedToken string, authCustomClaims *AuthCustomClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, authCustomClaims, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("DPzN3tMBaKsAPxvq8hWfaBHu5oeoj4bioNMQ6NzBSifkTthYAcoM67NzWTaZbPSDhGTkZhsdxyvYmNALanSoa3MH8CBW6Auv"), nil
	})

}

func existCoockie(coockie string, checker chan<- string) {
	var sessionID = SessionCredentials{}
	val := config.Client.Get(context.Background(), coockie).Val()
	json.Unmarshal([]byte(val), &sessionID)
	checker <- sessionID.ID
}
