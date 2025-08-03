package models

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	Base
	AlertID   uuid.UUID `json:"alert_id"`
	Alert     Alert     `gorm:"foreignKey:AlertID"`
	Status    string    `json:"status" gorm:"check:status IN ('open', 'dispatched', 'resolved')"`
	CreatedBy uuid.UUID `json:"created_by"`
	Operator  User      `gorm:"foreignKey:CreatedBy"`
	CreatedAt time.Time `json:"created_at"`
}
