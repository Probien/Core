package middleware

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/api/response"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/redis"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtRbac(authenticator *component.Authenticator, cookieManager *redis.Client, authorities ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := component.CustomClaims{}
		checker := make(chan bool, 1)
		cookie, err := c.Cookie("SID")
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{Status: http.StatusUnauthorized, Message: "Authorization is required", Data: "Cookie is not present in the request header"})
			return
		}

		authHeader := c.GetHeader("Authorization")
		go cookieManager.ExistCookie(cookie, checker)
		if len(authHeader) >= 0 && authHeader != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.Response{Status: http.StatusUnauthorized, Message: "Authorization is required", Data: "Token is not present in the request header"})
			return
		}

		splitToken := strings.Split(authHeader, "Bearer")
		encodedToken := strings.TrimSpace(splitToken[1])
		token, err := authenticator.ValidateAndParseToken(encodedToken, &data)
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, response.Response{Status: http.StatusUnauthorized, Message: "Authorization is required", Data: err})
			return
		}

		if authenticator.CheckAuthorities(authorities, &data) && <-checker {
			//extract user_id from parsed token for stored procedures
			userId, _ := strconv.Atoi(data.RegisteredClaims.Subject)
			//set user_id to request only for this context(request)
			c.Set("user_id", userId)
			c.Next()
		}

	}
}
