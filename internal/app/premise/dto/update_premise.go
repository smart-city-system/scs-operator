package dto

type UpdatePremiseDto struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Address string `json:"address" validate:"required,min=2,max=255"`
}
