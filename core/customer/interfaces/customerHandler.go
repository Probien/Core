package interfaces

import "github.com/gin-gonic/gin"

func CustomerHandler(v1 *gin.RouterGroup) {

	customerHandlerV1 := *v1.Group("/customer")

	customerHandlerV1.GET("", func(c *gin.Context) {

	})
}
