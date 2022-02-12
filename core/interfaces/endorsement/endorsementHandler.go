package endorsement_interface

import (
	"net/http"

	application "github.com/JairDavid/Probien-Backend/core/application/endorsement"
	"github.com/gin-gonic/gin"
)

func EndorsementHandler(v1 *gin.RouterGroup) {

	endorsementHandlerV1 := *v1.Group("/endorsement")
	interactor := application.EndorsemenInteractor{}

	endorsementHandlerV1.POST("/", func(c *gin.Context) {
		endorsement, err := interactor.Create(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": &endorsement})
	})

	endorsementHandlerV1.GET("/", func(c *gin.Context) {
		endorsements, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &endorsements})
		}
	})

	endorsementHandlerV1.GET("/:id", func(c *gin.Context) {
		endorsement, err := interactor.GetById(c)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &endorsement})
		}
	})

}
