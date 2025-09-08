package dto

type UpdateAlarmDto struct {
	Status string `json:"status" validate:"required"`
}
