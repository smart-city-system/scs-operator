package models

import "github.com/google/uuid"

type IncidentGuidance struct {
	Base
	IncidentID         uuid.UUID        `json:"incident_id"`
	Incident           Incident         `gorm:"foreignKey:IncidentID"`
	GuidanceTemplateID uuid.UUID        `json:"guidance_template_id"`
	GuidanceTemplate   GuidanceTemplate `gorm:"foreignKey:GuidanceTemplateID"`
	AssignedBy         uuid.UUID        `json:"assigned_by"`
	Operator           User             `gorm:"foreignKey:AssignedBy"`
	AssignedTo         uuid.UUID        `json:"assigned_to"`
	Guard              User             `gorm:"foreignKey:AssignedTo"`
}
