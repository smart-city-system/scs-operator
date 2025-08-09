package models

import (
	"github.com/google/uuid"
)

type Incident struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	AlarmID     uuid.UUID `json:"alarm_id,omitempty"`
	Alarm       *Alarm    `json:"alarm,omitempty" gorm:"foreignKey:AlarmID"`
	Status      string    `json:"status" gorm:"check:status IN ('new', 'in_progress', 'resolved')"`
}
