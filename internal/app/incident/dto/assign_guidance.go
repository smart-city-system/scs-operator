package dto

type AssignGuidance struct {
	GuidanceTemplateID string `json:"guidance_template_id" validate:"required"`
	Assignee           string `json:"assignee_id" validate:"required"`
}
