package dto

type CreatePremiseDto struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	Address         string `json:"address" validate:"required,min=2,max=255"`
	ParentPremiseID string `json:"parent_premise_id,omitempty" validate:"omitempty,uuid"`
}
