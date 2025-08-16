package services

import (
	"context"
	"scs-operator/internal/app/alarm/dto"
	alarmRepositories "scs-operator/internal/app/alarm/repository"
	premiseRepositories "scs-operator/internal/app/premise/repository"
	"scs-operator/internal/models"
	"scs-operator/pkg/errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	alarmRepo   alarmRepositories.AlarmRepository
	premiseRepo premiseRepositories.PremiseRepository
}

func NewAlarmService(alarmRepo alarmRepositories.AlarmRepository, premiseRepo premiseRepositories.PremiseRepository) *Service {
	return &Service{alarmRepo: alarmRepo, premiseRepo: premiseRepo}
}

func (s *Service) CreateAlarm(ctx context.Context, createAlarmDto *dto.CreateAlarmDto) (*models.Alarm, error) {
	alarm := &models.Alarm{
		Type:        createAlarmDto.Type,
		Description: createAlarmDto.Description,
		Severity:    createAlarmDto.Severity,
	}
	if createAlarmDto.TriggeredAt != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createAlarmDto.TriggeredAt)
		if err != nil {
			return nil, err
		}
		alarm.TriggeredAt = parsedTime
	}
	if createAlarmDto.PremiseID != "" {
		premiseID, err := uuid.Parse(createAlarmDto.PremiseID)

		if err != nil {
			return nil, err
		}

		premise, err := s.premiseRepo.GetPremiseByID(ctx, premiseID.String())

		if err != nil {
			return nil, err
		}
		alarm.Premise = premise
		alarm.PremiseID = premiseID
	}
	createdAlarm, err := s.alarmRepo.CreateAlarm(ctx, alarm)
	if err != nil {
		return nil, errors.NewDatabaseError("create alarm", err)
	}
	return createdAlarm, nil

}

func (s *Service) GetAlarms(ctx context.Context) ([]models.Alarm, error) {
	alarms, err := s.alarmRepo.GetAlarms(ctx)
	if err != nil {
		return nil, errors.NewDatabaseError("get alarms", err)
	}
	return alarms, nil
}
