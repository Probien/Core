package interfaces

import "github.com/gin-gonic/gin"

func EndorsementHandler(v1 *gin.RouterGroup) {

	endorsementHandlerV1 := *v1.Group("/endorsement")

	endorsementHandlerV1.GET("/", func(c *gin.Context) {

	})
}
