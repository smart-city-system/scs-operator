package dto

type CreateGuidanceTemplateDto struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	Steps       []Steps `json:"steps"`
}
type Steps struct {
	StepNumber  int    `json:"step_number" validate:"required"`
	Instruction string `json:"instruction" validate:"required"`
}
