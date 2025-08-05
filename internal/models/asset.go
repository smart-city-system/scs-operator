package models

import "github.com/google/uuid"

type Asset struct {
	Base
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	PremiseID uuid.UUID `json:"premise_id"`
	Premise   *Premise  `gorm:"foreignKey:PremiseID" json:"premise"`
	Type      string    `json:"type"`
	IsActive  bool      `json:"is_active"`
}
