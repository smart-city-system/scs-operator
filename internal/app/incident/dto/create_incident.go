package dto

type CreateIncidentDto struct {
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description" validate:"required"`
	AlarmId            string `json:"alarm_id" validate:"required"`
	Severity           string `json:"severity" validate:"required"`
	Location           string `json:"location" validate:"required"`
	GuidanceTemplateID string `json:"guidance_template_id" validate:"required"`
	Assignee           string `json:"assignee_id" validate:"required"`
}
