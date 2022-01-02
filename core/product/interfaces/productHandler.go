package interfaces

import "github.com/gin-gonic/gin"

func ProductHandler(v1 *gin.RouterGroup) {
	productHandlerV1 := *v1.Group("/product")
	productHandlerV1.GET("/", func(c *gin.Context) {

	})
}
