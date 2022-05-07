package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LogsRepositoryImp struct {
	database *gorm.DB
}

func NewLogsRepositoryImp(db *gorm.DB) repository.IlogsRepository {
	return &LogsRepositoryImp{database: db}
}

func (r *LogsRepositoryImp) GetAllSessions(c *gin.Context) (*[]domain.SessionLog, error) {
	return nil, nil
}

func (r *LogsRepositoryImp) GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, error) {
	return nil, nil
}

func (r *LogsRepositoryImp) GetAllPayments(c *gin.Context) (*[]domain.PaymentLog, error) {
	return nil, nil
}

func (r *LogsRepositoryImp) GetAllPaymentsByCustomerId(c *gin.Context) (*[]domain.PaymentLog, error) {
	return nil, nil
}

func (r *LogsRepositoryImp) GetAllMovements(c *gin.Context) (*[]domain.ModerationLog, error) {
	return nil, nil
}

func (r *LogsRepositoryImp) GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, error) {
	return nil, nil
}
