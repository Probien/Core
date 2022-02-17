package auth

import (
	"fmt"
	"net/http"

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
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
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
