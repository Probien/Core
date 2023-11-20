package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type LogApp struct {
	port port.ILogRepository
}

func NewLogApp(repository port.ILogRepository) LogApp {
	return LogApp{
		port: repository,
	}
}

func (l *LogApp) GetAllSessions(params url.Values) (*[]dto.SessionLog, map[string]interface{}, error) {
	return l.port.GetAllSessions(params)
}

func (l *LogApp) GetAllSessionsByEmployeeId(id int, params url.Values) (*[]dto.SessionLog, map[string]interface{}, error) {
	return l.port.GetAllSessionsByEmployeeId(id, params)
}

func (l *LogApp) GetAllMovements(params url.Values) (*[]dto.ModerationLog, map[string]interface{}, error) {
	return l.port.GetAllMovements(params)
}

func (l *LogApp) GetAllMovementsByEmployeeId(id int, params url.Values) (*[]dto.ModerationLog, map[string]interface{}, error) {
	return l.port.GetAllMovementsByEmployeeId(id, params)
}
