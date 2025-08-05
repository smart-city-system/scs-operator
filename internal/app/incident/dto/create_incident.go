package dto

type CreateIncidentDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	AssetID     string `json:"asset_id" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Status      string `json:"status" validate:"required"`
}
