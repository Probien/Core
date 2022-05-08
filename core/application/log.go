package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type LogsInteractor struct{}

func (li *LogsInteractor) GetAllSessions(c *gin.Context) (*[]domain.SessionLog, error) {
	repository := persistance.NewLogsRepositoryImp(config.Database)
	return repository.GetAllSessions(c)
}

func (li *LogsInteractor) GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, error) {
	repository := persistance.NewLogsRepositoryImp(config.Database)
	return repository.GetAllSessionsByEmployeeId(c)
}

func (li *LogsInteractor) GetAllPayments(c *gin.Context) (*[]domain.PaymentLog, error) {
	repository := persistance.NewLogsRepositoryImp(config.Database)
	return repository.GetAllPayments(c)
}

func (li *LogsInteractor) GetAllPaymentsByCustomerId(c *gin.Context) (*[]domain.PaymentLog, error) {
	repository := persistance.NewLogsRepositoryImp(config.Database)
	return repository.GetAllPaymentsByCustomerId(c)
}

func (li *LogsInteractor) GetAllMovements(c *gin.Context) (*[]domain.ModerationLog, error) {
	repository := persistance.NewLogsRepositoryImp(config.Database)
	return repository.GetAllMovements(c)
}

func (li *LogsInteractor) GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, error) {
	repository := persistance.NewLogsRepositoryImp(config.Database)
	return repository.GetAllMovementsByEmployeeId(c)
}
