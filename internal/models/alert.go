package models

import (
	"time"

	"github.com/google/uuid"
)

type Alert struct {
	Base
	AssetID     uuid.UUID
	Asset       Asset `gorm:"foreignKey:AssetID"`
	TriggeredAt time.Time
	Severity    string
	Description string
}
