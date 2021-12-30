package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryHandler(v1 *gin.RouterGroup) {

	categoryHandlerV1 := *v1.Group("/category")

	categoryHandlerV1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "world"})
	})
}
