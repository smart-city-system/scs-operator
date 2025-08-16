package dto

type CreateGuidanceTemplateDto struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Category    *string `json:"category"`
	Steps       []Steps `json:"steps"`
}
type Steps struct {
	StepNumber  int    `json:"step_number" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
