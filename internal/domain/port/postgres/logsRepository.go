package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type ILogRepository interface {
	GetAllSessions(params url.Values) (*[]dto.SessionLog, map[string]interface{}, error)
	GetAllSessionsByEmployeeId(id int, params url.Values) (*[]dto.SessionLog, map[string]interface{}, error)

	GetAllMovements(params url.Values) (*[]dto.ModerationLog, map[string]interface{}, error)
	GetAllMovementsByEmployeeId(id int, params url.Values) (*[]dto.ModerationLog, map[string]interface{}, error)
}
