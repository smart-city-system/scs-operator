package dto

type UpdateIncidentDto struct {
	Status string `json:"status" validate:"required"`
}
