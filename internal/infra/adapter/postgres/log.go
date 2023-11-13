package adapter

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter"
	"math"
	"net/url"

	"gorm.io/gorm"
)

type LogsRepositoryImp struct {
	database *gorm.DB
}

func NewLogsRepositoryImp(db *gorm.DB) port.ILogRepository {
	return &LogsRepositoryImp{database: db}
}

func (r *LogsRepositoryImp) GetAllSessions(params url.Values) (*[]dto.SessionLog, map[string]interface{}, error) {
	var sessions []dto.SessionLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("session_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.SessionLog{}).Scopes(adapter.Paginate(params, paginationResult)).Preload("Employee").Find(&sessions).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &sessions, paginationResult, nil
}

func (r *LogsRepositoryImp) GetAllSessionsByEmployeeId(id int, params url.Values) (*[]dto.SessionLog, map[string]interface{}, error) {
	var sessions []dto.SessionLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("session_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.SessionLog{}).Scopes(adapter.Paginate(params, paginationResult)).Where("employee_id = ?", id).Preload("Employee").Find(&sessions).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &sessions, paginationResult, nil
}

func (r *LogsRepositoryImp) GetAllMovements(params url.Values) (*[]dto.ModerationLog, map[string]interface{}, error) {
	var movements []dto.ModerationLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("moderation_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.ModerationLog{}).Scopes(adapter.Paginate(params, paginationResult)).Find(&movements).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &movements, paginationResult, nil
}

func (r *LogsRepositoryImp) GetAllMovementsByEmployeeId(id int, params url.Values) (*[]dto.ModerationLog, map[string]interface{}, error) {
	var movements []dto.ModerationLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	go r.database.Table("moderation_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.ModerationLog{}).Scopes(adapter.Paginate(params, paginationResult)).Where("user_id", id).Find(&movements).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &movements, paginationResult, nil
}
