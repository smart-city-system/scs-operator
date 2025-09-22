package types

import "scs-operator/internal/models"

// PremiseListResponse represents a paginated response for premises
type PremiseListResponse struct {
	Data       []models.Premise `json:"data"`
	Pagination Pagination       `json:"pagination"`
}

// IncidentListResponse represents a paginated response for incidents
type IncidentListResponse struct {
	Data       []models.Incident `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

// AlarmListResponse represents a response for alarms list
type AlarmListResponse []models.Alarm

// GuidanceTemplateListResponse represents a response for guidance templates list
type GuidanceTemplateListResponse []models.GuidanceTemplate

// GuidanceStepListResponse represents a response for guidance steps list
type GuidanceStepListResponse []models.GuidanceStep

// UserListResponse represents a response for users list
type UserListResponse []models.User
