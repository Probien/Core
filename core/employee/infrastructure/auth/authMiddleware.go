package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Prueba": "Se necesita un token"})
		c.Abort()
	}
}
