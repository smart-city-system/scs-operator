package dto

type CreateIncidentDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	AlarmId     string `json:"alarm_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Severity    string `json:"severity" validate:"required"`
	Location    string `json:"location" validate:"required"`
}
