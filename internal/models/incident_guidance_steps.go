package models

import (
	"time"

	"github.com/google/uuid"
)

type IncidentGuidanceStep struct {
	Base
	IncidentGuidanceID uuid.UUID        `json:"incident_guidance_id"`
	IncidentGuidance   IncidentGuidance `gorm:"foreignKey:IncidentGuidanceID"`
	StepNumber         int64            `json:"step_number"`
	IsCompleted        bool             `json:"is_completed" gorm:"default:false"`
	CompletedAt        *time.Time       `json:"completed_at,omitempty"`
}
