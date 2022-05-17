package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		data := AuthCustomClaims{}

		if len(authHeader) > 0 && authHeader != "Bearer" {

			splitToken := strings.Split(authHeader, "Bearer")
			encodedToken := strings.TrimSpace(splitToken[1])
			token, err := validateAndParseToken(encodedToken, &data)
			log.Print(err)

			if token.Valid {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, common.Response{Status: 500, Message: "Authorization is required", Data: "Invalid token"})
				c.AbortWithStatus(http.StatusUnauthorized)
			}

		} else {
			c.JSON(http.StatusUnauthorized, common.Response{Status: 500, Message: "Authorization is required", Data: "Unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func RoutesAndAuthority(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		data := AuthCustomClaims{}

		if len(authHeader) > 0 && authHeader != "Bearer" {
			splitToken := strings.Split(authHeader, "Bearer")
			encodedToken := strings.TrimSpace(splitToken[1])
			validateAndParseToken(encodedToken, &data)

			if data.IsAdmin == isAdmin {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, common.Response{Status: 500, Message: "Authorization is required", Data: "Unauthorized"})
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.JSON(http.StatusUnauthorized, common.Response{Status: 500, Message: "Authorization is required", Data: "Unauthorized"})
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
