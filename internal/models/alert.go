package models

import (
	"time"

	"github.com/google/uuid"
)

type Alert struct {
	Base
	CameraID    uuid.UUID
	Camera      Camera `gorm:"foreignKey:CameraID"`
	TriggeredAt time.Time
	Severity    string
	Description string
}
