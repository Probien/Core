package persistence

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
	var sessions []domain.SessionLog

	if err := r.database.Model(&domain.SessionLog{}).Preload("Employee").Find(&sessions).Error; err != nil {
		return nil, ErrorProcess
	}

	return &sessions, nil
}

func (r *LogsRepositoryImp) GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, error) {
	var sessions []domain.SessionLog

	if err := r.database.Model(&domain.SessionLog{}).Where("employee_id = ?", c.Param("id")).Preload("Employee").Find(&sessions).Error; err != nil {
		return nil, ErrorProcess
	}

	return &sessions, nil
}

func (r *LogsRepositoryImp) GetAllPayments(c *gin.Context) (*[]domain.PaymentLog, error) {
	var payments []domain.PaymentLog

	if err := r.database.Model(&domain.PaymentLog{}).Preload("Employee").Preload("Customer").Find(&payments).Error; err != nil {
		return nil, ErrorProcess
	}

	return &payments, nil
}

func (r *LogsRepositoryImp) GetAllPaymentsByCustomerId(c *gin.Context) (*[]domain.PaymentLog, error) {
	var payments []domain.PaymentLog

	if err := r.database.Model(&domain.PaymentLog{}).Where("customer_id = ?", c.Param("id")).Preload("Employee").Preload("Customer").Find(&payments).Error; err != nil {
		return nil, ErrorProcess
	}

	return &payments, nil
}

func (r *LogsRepositoryImp) GetAllMovements(c *gin.Context) (*[]domain.ModerationLog, error) {
	var movements []domain.ModerationLog

	if err := r.database.Model(&domain.ModerationLog{}).Find(&movements).Error; err != nil {
		return nil, ErrorProcess
	}

	return &movements, nil
}

func (r *LogsRepositoryImp) GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, error) {
	var movements []domain.ModerationLog

	if err := r.database.Model(&domain.ModerationLog{}).Where("user_id", c.Param("id")).Find(&movements).Error; err != nil {
		return nil, ErrorProcess
	}

	return &movements, nil
}
