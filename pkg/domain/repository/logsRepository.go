package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/pkg/domain"
)

type IlogsRepository interface {
	GetAllSessions(params url.Values) (*[]domain.SessionLog, map[string]interface{}, error)
	GetAllSessionsByEmployeeId(id int, params url.Values) (*[]domain.SessionLog, map[string]interface{}, error)

	GetAllMovements(params url.Values) (*[]domain.ModerationLog, map[string]interface{}, error)
	GetAllMovementsByEmployeeId(id int, params url.Values) (*[]domain.ModerationLog, map[string]interface{}, error)
}
