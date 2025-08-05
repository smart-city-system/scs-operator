package dto

type CreateAssetDto struct {
	Name      string `json:"name" validate:"required"`
	Location  string `json:"location" validate:"required"`
	PremiseID string `json:"premise_id" validate:"required"`
	Type      string `json:"type" validate:"required"`
	IsActive  bool   `json:"is_active"`
}
