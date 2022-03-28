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
			ts := authHeader[len("Bearer "):]
			token, err := validateToken(ts)
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
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

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("DPzN3tMBaKsAPxvq8hWfaBHu5oeoj4bioNMQ6NzBSifkTthYAcoM67NzWTaZbPSDhGTkZhsdxyvYmNALanSoa3MH8CBW6Auv"), nil
	})

}
