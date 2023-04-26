package auth

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func JwtRbac(authorities ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		checker := make(chan bool, 1)
		data := CustomClaims{}
		cookie, _ := c.Cookie("SID")
		authHeader := c.GetHeader("Authorization")
		go existCookie(cookie, checker)

		if len(authHeader) > 0 && authHeader != "Bearer" {
			splitToken := strings.Split(authHeader, "Bearer")
			encodedToken := strings.TrimSpace(splitToken[1])
			token, _ := validateAndParseToken(encodedToken, &data)

			if token.Valid && checkAuthorities(authorities, &data) && <-checker {
				//extract user_id from parsed token for stored procedures
				userId, _ := strconv.Atoi(data.RegisteredClaims.Subject)
				//set user_id to request only for this context(request)
				c.Set("user_id", userId)
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
