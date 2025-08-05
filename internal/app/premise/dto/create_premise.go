package dto

type CreatePremiseDto struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	Location        string `json:"location" validate:"required,min=2,max=255"`
	ParentPremiseID string `json:"parent_premise_id,omitempty" validate:"omitempty,uuid"`
}
