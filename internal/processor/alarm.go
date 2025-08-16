package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"scs-operator/internal/app/alarm/dto"
	services "scs-operator/internal/app/alarm/service"
	"scs-operator/pkg/logger"

	"github.com/segmentio/kafka-go"
)

type Processor interface {
	Process(msg kafka.Message) error
}

type AlarmProcessor struct {
	alarmService services.Service
	logger       logger.Logger
}

func NewAlarmProcessor(alarmService services.Service, logger logger.Logger) Processor {
	return &AlarmProcessor{alarmService: alarmService, logger: logger}
}

func (ap AlarmProcessor) Process(msg kafka.Message) error {
	// Simulate processing logic
	fmt.Printf("Processing ")
	var createAlarmDto dto.CreateAlarmDto
	err := json.Unmarshal(msg.Value, &createAlarmDto)
	if err != nil {
		ap.logger.Errorf("Failed to unmarshal message: %v", err)
		return err
	}
	_, err = ap.alarmService.CreateAlarm(context.Background(), &createAlarmDto)
	if err != nil {
		ap.logger.Errorf("Failed to create alarm: %v", err)
		return err
	}
	ap.logger.Info("Alarm created")
	return nil // or return an actual error if something goes wrong
}
