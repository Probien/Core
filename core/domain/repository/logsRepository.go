package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type IlogsRepository interface {
	GetAllSessions(c *gin.Context) (*[]domain.SessionLog, error)
	GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, error)

	GetAllPayments(c *gin.Context) (*[]domain.PaymentLog, error)
	GetAllPaymentsByCustomerId(c *gin.Context) (*[]domain.PaymentLog, error)

	GetAllMovements(c *gin.Context) (*[]domain.ModerationLog, error)
	GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, error)
}
