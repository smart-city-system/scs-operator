package dto

type UpdateGuidanceTemplateDto struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Category    *string  `json:"category"`
	AddSteps    []Step   `json:"add_steps"`
	UpdateSteps []Step   `json:"update_steps"`
	RemoveSteps []string `json:"remove_steps"`
}
