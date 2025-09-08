package services

import (
	"context"
	"encoding/json"
	"scs-operator/internal/app/alarm/dto"
	alarmRepositories "scs-operator/internal/app/alarm/repository"
	premiseRepositories "scs-operator/internal/app/premise/repository"
	"scs-operator/internal/models"
	"scs-operator/pkg/errors"
	kafka_client "scs-operator/pkg/kafka"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Service struct {
	alarmRepo   alarmRepositories.AlarmRepository
	premiseRepo premiseRepositories.PremiseRepository
	producer    kafka_client.Producer
}

func NewAlarmService(alarmRepo alarmRepositories.AlarmRepository, premiseRepo premiseRepositories.PremiseRepository, producer kafka_client.Producer) *Service {
	return &Service{alarmRepo: alarmRepo, premiseRepo: premiseRepo, producer: producer}
}

func (s *Service) CreateAlarm(ctx context.Context, createAlarmDto *dto.CreateAlarmDto) (*models.Alarm, error) {
	alarm := &models.Alarm{
		Type:        createAlarmDto.Type,
		Description: createAlarmDto.Description,
		Severity:    createAlarmDto.Severity,
		Status:      "new",
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

	// Send Kafka message after successful alarm creation
	alarmData, err := json.Marshal(createdAlarm)
	if err != nil {
		// Log error but don't fail the operation since alarm was created successfully
		// You might want to add proper logging here
		return createdAlarm, nil
	}

	producerMessage := kafka.Message{
		Key:   []byte(createdAlarm.ID.String()),
		Value: alarmData,
	}

	if err := s.producer.WriteMessages(ctx, producerMessage); err != nil {
		// Log error but don't fail the operation since alarm was created successfully
		// You might want to add proper logging here
		return createdAlarm, nil
	}

	return createdAlarm, nil

}

func (s *Service) GetAlarms(ctx context.Context, status string) ([]models.Alarm, error) {
	alarms, err := s.alarmRepo.GetAlarms(ctx, status)
	if err != nil {
		return nil, errors.NewDatabaseError("get alarms", err)
	}
	return alarms, nil
}

func (s *Service) UpdateAlarm(ctx context.Context, id string, updateAlarmDto *dto.UpdateAlarmDto) (*models.Alarm, error) {
	alarm, err := s.alarmRepo.GetAlarmByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("alarm not found")
	}
	alarm.Status = updateAlarmDto.Status
	updatedAlarm, err := s.alarmRepo.UpdateAlarm(ctx, id, alarm)
	if err != nil {
		return nil, errors.NewDatabaseError("update alarm", err)
	}
	return updatedAlarm, nil
}
