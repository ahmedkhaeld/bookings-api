package models

import "time"

// Reservation is the reservation model
type Reservation struct {
	ID        int       `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	RoomID    int       `json:"room_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Room      Room      `json:"room"`
	Processed int       `json:"processed,omitempty"`
}
