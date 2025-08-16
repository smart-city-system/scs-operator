package dto

type CreateAlarmDto struct {
	PremiseID   string `json:"premise_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	TriggeredAt string `json:"triggered_at"`
	Severity    string `json:"severity"`
}
