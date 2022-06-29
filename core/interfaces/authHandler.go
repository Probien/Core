package interfaces

import (
	"log"
	"net/http"
	"strings"

	"context"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func AuthHandler(v1 *gin.RouterGroup) {

	security := *v1.Group("/auth")
	interactor := application.EmployeeInteractor{}

	security.POST("/login", func(c *gin.Context) {
		tokenizer := make(chan string, 1)
		employee, err := interactor.Login(c)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			interactor.GenerateToken(employee, tokenizer)
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employee, Token: <-tokenizer})
		}
	})

	security.POST("/logout", func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		splitToken := strings.Split(header, "Bearer")
		encodedToken := strings.TrimSpace(splitToken[1])
		_, err := config.Client.Get(ctx, encodedToken).Result()
		if err == redis.Nil {
			log.Print(err)
		}
	})

}
