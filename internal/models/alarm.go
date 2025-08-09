package models

import "github.com/google/uuid"

// Alarm represents an alarm in the SCS system.
type Alarm struct {
	ID          string    `json:"id"`
	PremiseID   uuid.UUID `json:"premise_id"`
	Premise     *Premise  `json:"premise,omitempty" gorm:"foreignkey:PremiseID"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	TriggeredAt string    `json:"triggered_at"`
}
