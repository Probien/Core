package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type IlogsRepository interface {
	GetAllSessions() (*[]domain.SessionLog, error)
	GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, error)

	GetAllPayments() (*[]domain.PaymentLog, error)
	GetAllPaymentsByCustomerId(c *gin.Context) (*[]domain.PaymentLog, error)

	GetAllMovements() (*[]domain.ModerationLog, error)
	GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, error)
}
