package application

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"net/url"
)

type LogsInteractor struct {
	repository repository.IlogsRepository
}

func NewLogsInteractor(repository repository.IlogsRepository) LogsInteractor {
	return LogsInteractor{
		repository: repository,
	}
}

func (l *LogsInteractor) GetAllSessions(params url.Values) (*[]domain.SessionLog, map[string]interface{}, error) {
	return l.repository.GetAllSessions(params)
}

func (l *LogsInteractor) GetAllSessionsByEmployeeId(id int, params url.Values) (*[]domain.SessionLog, map[string]interface{}, error) {
	return l.repository.GetAllSessionsByEmployeeId(id, params)
}

func (l *LogsInteractor) GetAllMovements(params url.Values) (*[]domain.ModerationLog, map[string]interface{}, error) {
	return l.repository.GetAllMovements(params)
}

func (l *LogsInteractor) GetAllMovementsByEmployeeId(id int, params url.Values) (*[]domain.ModerationLog, map[string]interface{}, error) {
	return l.repository.GetAllMovementsByEmployeeId(id, params)
}
