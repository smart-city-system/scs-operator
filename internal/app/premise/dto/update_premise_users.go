package dto

type UpdatePremiseUserDto struct {
	AddedUsers   []string `json:"added_users" validate:"omitempty"`
	RemovedUsers []string `json:"removed_users" validate:"omitempty"`
}
