package dto

type CreateCameraDto struct {
	Name      string `json:"name" validate:"required"`
	Location  string `json:"location" validate:"required"`
	PremiseID string `json:"premise_id" validate:"required"`
	StreamURL string `json:"stream_url" validate:"required"`
	IsActive  bool   `json:"is_active"`
}
