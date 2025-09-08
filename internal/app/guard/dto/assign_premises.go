package dto

type AssignPremisesDto struct {
	GuardID   string `json:"guard_id" validate:"required"`
	PremiseID string `json:"premise_id" validate:"required"`
}
