package models

import (
	"time"

	"github.com/google/uuid"
)

type DispatchRequest struct {
	Base
	IncidentID   uuid.UUID `json:"incident_id"`
	Incident     Incident  `gorm:"foreignKey:IncidentID"`
	GuardID      uuid.UUID `json:"guard_id"`
	Guard        User      `gorm:"foreignKey:GuardID"`
	Instruction  string    `json:"instruction"`
	Status       string    `gorm:"default:pending;check:status IN ('pending', 'acknowledged', 'completed')"`
	DispatchedAt time.Time `json:"dispatched_at"`
}
