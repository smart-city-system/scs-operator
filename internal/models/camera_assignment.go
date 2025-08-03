package models

import (
	"time"

	"github.com/google/uuid"
)

type CameraAssignment struct {
	Base
	GuardID    uuid.UUID `json:"guard_id"`
	Guard      User      `gorm:"foreignKey:GuardID"`
	CameraID   uuid.UUID `json:"camera_id"`
	Camera     Camera    `gorm:"foreignKey:CameraID"`
	AssignedAt time.Time `json:"assigned_at"`
}
