package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type LogsInteractor struct{}

func (li *LogsInteractor) GetAllSessions(params url.Values) (*[]domain.SessionLog, map[string]interface{}, error) {
	repository := postgres.NewLogsRepositoryImp(config.Database)
	return repository.GetAllSessions(params)
}

func (li *LogsInteractor) GetAllSessionsByEmployeeId(id int) (*[]domain.SessionLog, map[string]interface{}, error) {
	repository := postgres.NewLogsRepositoryImp(config.Database)
	return repository.GetAllSessionsByEmployeeId(id)
}

func (li *LogsInteractor) GetAllMovements(params url.Values) (*[]domain.ModerationLog, map[string]interface{}, error) {
	repository := postgres.NewLogsRepositoryImp(config.Database)
	return repository.GetAllMovements(params)
}

func (li *LogsInteractor) GetAllMovementsByEmployeeId(id int) (*[]domain.ModerationLog, map[string]interface{}, error) {
	repository := postgres.NewLogsRepositoryImp(config.Database)
	return repository.GetAllMovementsByEmployeeId(id)
}
