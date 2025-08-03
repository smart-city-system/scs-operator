package dto

type CreatePremiseDto struct {
	Name            string `json:"name" validate:"required"`
	Location        string `json:"location" validate:"required"`
	ParentPremiseID string `json:"parent_premise_id,omitempty"`
}
