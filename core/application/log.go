package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type LogsInteractor struct{}

func (li *LogsInteractor) GetAllSessions() (*[]domain.SessionLog, error) {
	repository := persistence.NewLogsRepositoryImp(config.Database)
	return repository.GetAllSessions()
}

func (li *LogsInteractor) GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, error) {
	repository := persistence.NewLogsRepositoryImp(config.Database)
	return repository.GetAllSessionsByEmployeeId(c)
}

func (li *LogsInteractor) GetAllPayments() (*[]domain.PaymentLog, error) {
	repository := persistence.NewLogsRepositoryImp(config.Database)
	return repository.GetAllPayments()
}

func (li *LogsInteractor) GetAllPaymentsByCustomerId(c *gin.Context) (*[]domain.PaymentLog, error) {
	repository := persistence.NewLogsRepositoryImp(config.Database)
	return repository.GetAllPaymentsByCustomerId(c)
}

func (li *LogsInteractor) GetAllMovements() (*[]domain.ModerationLog, error) {
	repository := persistence.NewLogsRepositoryImp(config.Database)
	return repository.GetAllMovements()
}

func (li *LogsInteractor) GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, error) {
	repository := persistence.NewLogsRepositoryImp(config.Database)
	return repository.GetAllMovementsByEmployeeId(c)
}
