package repositories

import (
	"context"
	"fmt"
	"scs-operator/internal/models"

	"gorm.io/gorm"
)

type AlarmRepository struct {
	db *gorm.DB
}

func NewAlarmRepository(db *gorm.DB) *AlarmRepository {
	return &AlarmRepository{db: db}
}

func (r *AlarmRepository) CreateAlarm(ctx context.Context, Alarm *models.Alarm) (*models.Alarm, error) {
	if err := r.db.WithContext(ctx).Create(Alarm).Error; err != nil {
		return nil, fmt.Errorf("failed to create Alarm: %w", err)
	}
	return Alarm, nil
}
func (r *AlarmRepository) GetAlarms(ctx context.Context, status string) ([]models.Alarm, error) {
	var Alarms []models.Alarm
	query := r.db.WithContext(ctx).Preload("Premise")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&Alarms).Error; err != nil {
		return nil, fmt.Errorf("failed to get Alarms: %w", err)
	}
	return Alarms, nil
}

func (r *AlarmRepository) GetAlarmByID(ctx context.Context, id string) (*models.Alarm, error) {
	var Alarm models.Alarm

	if err := r.db.WithContext(ctx).First(&Alarm, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get Alarm: %w", err)
	}

	return &Alarm, nil
}

func (r *AlarmRepository) UpdateAlarm(ctx context.Context, id string, Alarm *models.Alarm) (*models.Alarm, error) {
	result := r.db.WithContext(ctx).Model(&models.Alarm{}).Where("id = ?", id).Updates(Alarm)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update Alarm: %w", result.Error)
	}
	return Alarm, nil
}
