package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func JwtAuth(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		data := AuthCustomClaims{}

		if len(authHeader) > 0 && authHeader != "Bearer" {
			splitToken := strings.Split(authHeader, "Bearer")
			encodedToken := strings.TrimSpace(splitToken[1])
			token, _ := validateAndParseToken(encodedToken, &data)

			if token.Valid && data.IsAdmin == isAdmin {
				user_id, _ := strconv.Atoi(data.RegisteredClaims.Subject)
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
