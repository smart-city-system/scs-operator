package models

import "github.com/google/uuid"

type Premise struct {
	Base
	Name            string     `json:"name"`
	Location        string     `json:"location"`
	ParentPremiseID *uuid.UUID `json:"parent_premise_id,omitempty"`
	ParentPremise   *Premise   `gorm:"foreignKey:ParentPremiseID"`
}
