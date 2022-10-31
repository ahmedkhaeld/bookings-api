package models

type AvailabilityPayload struct {
	StartDate string                 `json:"startDate"`
	EndDate   string                 `json:"endDate"`
	RoomID    string                 `json:"roomID,omitempty"`
	Message   string                 `json:"Message"`
	Ok        bool                   `json:"ok"`
	Rooms     map[string]interface{} `json:"rooms,omitempty"`
}

type ReservationPayload struct {
	ID        string                 `json:"id,omitempty"`
	FistName  string                 `json:"firstName"`
	LastName  string                 `json:"lastName"`
	Email     string                 `json:"email"`
	Phone     string                 `json:"phone"`
	StartDate string                 `json:"startDate"`
	EndDate   string                 `json:"endDate"`
	RoomID    string                 `json:"roomID,omitempty"`
	Room      map[string]interface{} `json:"room,omitempty"`
	Message   string                 `json:"Message"`
	Ok        bool                   `json:"ok"`
}
