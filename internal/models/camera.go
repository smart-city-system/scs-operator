package models

import "github.com/google/uuid"

type Camera struct {
	Base
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	PremiseID uuid.UUID `json:"premise_id"`
	Premise   *Premise  `gorm:"foreignKey:PremiseID" json:"premise"`
	StreamURL string    `json:"stream_url"`
	IsActive  bool      `json:"is_active"`
}
