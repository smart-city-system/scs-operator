package models

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	AlertID     *uuid.UUID `json:"alert_id,omitempty"`
	Alert       *Alert     `json:"alert,omitempty" gorm:"foreignKey:AlertID"`
	AssetID     uuid.UUID  `json:"asset_id"`
	Asset       *Asset     `json:"asset,omitempty" gorm:"foreignKey:AssetID"`
	Location    string     `json:"location"`
	Status      string     `json:"status" gorm:"check:status IN ('new', 'in_progress', 'resolved')"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty"`
	Operator    *User      `json:"operator,omitempty" gorm:"foreignKey:CreatedBy"`
	CreatedAt   time.Time  `json:"created_at"`
}
