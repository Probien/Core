package auth

import (
	"fmt"
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 0 {
			encodedToken := authHeader[len("Bearer "):]
			data := AuthCustomClaims{}
			token, err := validateAndParseToken(encodedToken, &data)
			if token.Valid {
				c.Next()
			} else {
				fmt.Println(err)
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
		encodedToken := authHeader[len("Bearer "):]
		data := AuthCustomClaims{}

		validateAndParseToken(encodedToken, &data)

		if data.IsAdmin == isAdmin {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, common.Response{Status: 500, Message: "Authorization is required", Data: "Unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func validateAndParseToken(encodedToken string, authCustomClaims *AuthCustomClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, authCustomClaims, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("DPzN3tMBaKsAPxvq8hWfaBHu5oeoj4bioNMQ6NzBSifkTthYAcoM67NzWTaZbPSDhGTkZhsdxyvYmNALanSoa3MH8CBW6Auv"), nil
	})

}
