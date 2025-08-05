package models

import (
	"time"

	"github.com/google/uuid"
)

type AssetAssignment struct {
	Base
	GuardID    uuid.UUID `json:"guard_id"`
	Guard      User      `gorm:"foreignKey:GuardID"`
	AssetID    uuid.UUID `json:"asset_id"`
	Asset      Asset     `gorm:"foreignKey:AssetID"`
	AssignedAt time.Time `json:"assigned_at"`
}
