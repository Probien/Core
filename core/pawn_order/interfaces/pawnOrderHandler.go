package interfaces

import "github.com/gin-gonic/gin"

func PawnOrderHandler(v1 *gin.RouterGroup) {

	pawnOrderHandlerV1 := *v1.Group("/pawn-order")

	pawnOrderHandlerV1.GET("/", func(c *gin.Context) {

	})
}
