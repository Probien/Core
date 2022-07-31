package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	roles := make(map[string]string)

	for k, v := range employee.Roles {
		roles["role_"+strconv.Itoa(k)] = v.Role.RoleName
	}

	claims := &AuthCustomClaims{
		Name:    employee.Profile.Name,
		IsAdmin: employee.IsAdmin,
		Roles:   roles,
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

func JwtAuth(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		checker := make(chan bool, 1)
		data := AuthCustomClaims{}
		coockie, _ := c.Cookie("SID")
		authHeader := c.GetHeader("Authorization")
		go existCoockie(coockie, checker)

		if len(authHeader) > 0 && authHeader != "Bearer" {
			splitToken := strings.Split(authHeader, "Bearer")
			encodedToken := strings.TrimSpace(splitToken[1])
			token, _ := validateAndParseToken(encodedToken, &data)

			if token.Valid && data.IsAdmin == isAdmin && <-checker {

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

		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

}

func existCoockie(coockie string, checker chan<- bool) {
	var sessionID = SessionCredentials{}
	val := config.Client.Get(context.Background(), coockie).Val()
	json.Unmarshal([]byte(val), &sessionID)
	checker <- val != "" && coockie == sessionID.ID

}
