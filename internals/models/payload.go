package models

type AvailabilityPayload struct {
	StartDate string                 `json:"startDate"`
	EndDate   string                 `json:"endDate"`
	RoomID    string                 `json:"roomID,omitempty"`
	Message   string                 `json:"Message"`
	Ok        bool                   `json:"ok"`
	Rooms     map[string]interface{} `json:"rooms,omitempty"`
}
