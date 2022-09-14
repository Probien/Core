package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type IlogsRepository interface {
	GetAllSessions(c *gin.Context) (*[]domain.SessionLog, map[string]interface{}, error)
	GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, map[string]interface{}, error)

	GetAllMovements(c *gin.Context) (*[]domain.ModerationLog, map[string]interface{}, error)
	GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, map[string]interface{}, error)
}
