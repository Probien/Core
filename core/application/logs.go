package application

import (
	"github.com/JairDavid/Probien-Backend/config/migrations/models"
	"github.com/gin-gonic/gin"
)

type LogsInteractor struct{}

func (li *LogsInteractor) GetAllSessions(c *gin.Context) (*[]models.SessionLog, error) {
	return nil, nil
}
