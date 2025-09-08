package models

import "github.com/google/uuid"

type GuardPremise struct {
	Base
	GuardID   uuid.UUID `json:"guard_id"`
	Guard     *User     `json:"guard,omitempty" gorm:"foreignKey:GuardID"`
	PremiseID uuid.UUID `json:"premise_id"`
	Premise   *Premise  `json:"premise,omitempty" gorm:"foreignKey:PremiseID"`
}