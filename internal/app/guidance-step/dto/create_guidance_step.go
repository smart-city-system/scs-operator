package dto

type CreateGuidanceStepDto struct {
	GuidanceTemplateID string `json:"guidance_template_id" validate:"required"`
	StepNumber         int    `json:"step_number" validate:"required"`
	Instruction        string `json:"instruction" validate:"required"`
}
